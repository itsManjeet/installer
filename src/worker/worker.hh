#ifndef WORKER_HH
#define WORKER_HH

#include <gtkmm.h>

#include <mutex>
#include <thread>

#include "../internal/data/data.hh"
#include "process.hh"
class ProcessPage;

class Worker {
 protected:
  mutable std::mutex m_Mutex;
  double m_Progress;
  Glib::ustring m_Message, m_Title;
  bool m_Success = false;
  bool m_Running = false;

  bool m_ExitOnFail = false;

  InstallData* m_InstallData;

  std::vector<std::shared_ptr<Process>> m_Processes;

 public:
  Worker(InstallData* data, std::vector<std::shared_ptr<Process>> processes,
         bool exitOnFail = false)
      : m_InstallData(data), m_Processes(processes), m_ExitOnFail(exitOnFail) {}

  void start(ProcessPage* page);

  void get(double* progress, Glib::ustring* mesg, Glib::ustring* title,
           bool* success) const;

  bool stopped() const { return !m_Running; }
  bool success() const { return m_Success; }
};

#endif