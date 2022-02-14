#ifndef PROCESS_HH
#define PROCESS_HH

#include <memory>
#include <string>
#include <vector>

#include "../internal/data/data.hh"
#include "../locale.hh"

class Process {
 protected:
  std::string m_Mesg;
  std::string m_Title;
  InstallData* m_Data;

 public:
  Process(std::string title, InstallData* data)
      : m_Data(data), m_Title(title) {}
  std::string const& title() const { return m_Title; }
  std::string const& mesg() const { return m_Mesg; }
  virtual bool process() = 0;
};

#endif