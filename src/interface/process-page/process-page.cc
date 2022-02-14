#include "process-page.hh"

#include "../../worker/worker.hh"

ProcessPage::ProcessPage(std::string const& title,
                         std::vector<std::shared_ptr<Process>> processes,
                         InstallData* data, bool exitOnError)
    : Page(data),
      m_WorkerThread(nullptr),
      m_Worker(data, processes, exitOnError) {
  m_Title.set_markup("<span size=\"xx-large\" weight=\"ultrabold\">" + title +
                     "</span>");
  m_Title.set_margin_top(27);
  m_Title.set_margin_bottom(17);
  pack_start(m_Title, false, false);
  pack_start(m_ListBox, true, true, 10);

  m_ProgressBar.set_show_text(true);
  m_ProgressBar.set_valign(Gtk::ALIGN_END);
  pack_start(m_ProgressBar, false, false);

  m_Dispatcher.connect(sigc::mem_fun(*this, &ProcessPage::on_notify));
}

void ProcessPage::on_load() {
  if (!m_WorkerThread && !alreadyDone) {
    isRunning = true;
    m_WorkerThread = new std::thread([this]() { m_Worker.start(this); });
    m_signal_update_buttons.emit(false, false);
  }
}

void ProcessPage::on_notify() {
  if (m_WorkerThread && m_Worker.stopped()) {
    if (m_WorkerThread->joinable()) {
      m_WorkerThread->join();
    }

    delete m_WorkerThread;
    m_WorkerThread = nullptr;
  }

  update();
}

void ProcessPage::update() {
  double progress;
  Glib::ustring message, title;
  bool success;

  m_Worker.get(&progress, &message, &title, &success);
  std::string img = "dialog-ok";
  if (!success) {
    img = "dialog-error";
  }

  auto button = ListButton<int>::create(img, title, message, 1);
  button->show_all();

  m_ListBox.append(button.get());
  m_ProgressBar.set_fraction(progress);
  m_ProgressBar.set_text(message);

  if (m_Worker.stopped()) {
    alreadyDone = true;
    if (success) {
      m_signal_update_buttons.emit(true, true);
    } else {
      m_signal_update_buttons.emit(false, false);
    }
  }
}

void ProcessPage::notify() { m_Dispatcher.emit(); }