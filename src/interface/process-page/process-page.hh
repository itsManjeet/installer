#ifndef PROCESS_PAGE_HH
#define PROCESS_PAGE_HH

#include <thread>

#include "../../worker/worker.hh"
#include "../listbox/listbox.hh"
#include "../page.hh"

class ProcessPage : public Page {
 protected:
  Gtk::Label m_Title;
  ListBox m_ListBox;
  GtkWidget* m_Clamp;
  Gtk::ProgressBar m_ProgressBar;

  bool isRunning = false;
  bool isError = false;
  bool alreadyDone = false;

  Worker m_Worker;

  Glib::Dispatcher m_Dispatcher;
  std::thread* m_WorkerThread;

 public:
  ProcessPage(std::string const& title,
              std::vector<std::shared_ptr<Process>> processes,
              InstallData* data,
              bool exitOnError = false);

  ~ProcessPage() {}
  void on_load();
  void on_notify();
  void update();
  void notify();
};

#endif