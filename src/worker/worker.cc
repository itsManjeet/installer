#include "worker.hh"

#include "../interface/process-page/process-page.hh"

void Worker::get(double* progress, Glib::ustring* mesg, Glib::ustring* title,
                 bool* success) const {
  std::lock_guard<std::mutex> lock(m_Mutex);
  if (progress) {
    *progress = m_Progress;
  }
  if (mesg) {
    *mesg = m_Message;
  }

  if (success) {
    *success = m_Success;
  }

  if (title) {
    *title = m_Title;
  }
}

void Worker::start(ProcessPage* page) {
  double progress = 0.0;
  double increment = 1.0 / m_Processes.size();
  bool all_pass = true;
  for (auto const& i : m_Processes) {
    bool status = i->process();
    {
      std::lock_guard<std::mutex> lock(m_Mutex);
      m_Message = i->mesg();
      m_Title = i->title();
      m_Progress += increment;
      m_Running = true;
      if (status) {
        m_Success = true;
      } else {
        m_Success = false;
        all_pass = false;
      }
      page->notify();
    }
    usleep(300000);

    if (m_ExitOnFail && !status) {
      std::lock_guard<std::mutex> lock(m_Mutex);
      m_Running = false;
      all_pass = false;
      break;
    }
  }

  {
    std::lock_guard<std::mutex> lock(m_Mutex);
    m_Title = "Result";
    m_Progress = 1.0;
    if (all_pass) {
      m_Message = "process complete successfully";
    } else {
      m_Message = "one or more process failed";
    }

    m_Success = all_pass;
    m_Running = false;
    page->notify();
  }
}