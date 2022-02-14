#ifndef APPLICATION_HH
#define APPLICATION_HH

#include <gtkmm.h>

#include "window.hh"

class Application : public Gtk::Application {
 private:
  Application();
  Window* create_window();

  void on_activate() override;
  void on_startup() override;

  void on_hide_window(Gtk::Window* window);

 public:
  virtual ~Application() override;

  static Glib::RefPtr<Application> create();
};

#endif