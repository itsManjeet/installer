#ifndef CHECKUP_HH
#define CHECKUP_HH

#include <memory>
#include <string>
#include <vector>

#include "../../locale.hh"
#include "../../worker/worker.hh"
#include "../data/data.hh"

class Checkup {
 public:
  static std::vector<std::shared_ptr<Process>> list(InstallData* data);
};

#endif