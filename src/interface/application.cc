#include "application.hh"

#include <iostream>

#include "config.h"

Application::Application() : Gtk::Application(APP_ID) {}
Application::~Application() {}

Glib::RefPtr<Application> Application::create() {
  return Glib::RefPtr<Application>(new Application());
}

Window* Application::create_window() {
  auto window = Window::create();
  add_window(*window);
  window->signal_hide().connect(
      sigc::bind(sigc::mem_fun(*this, &Application::on_hide_window), window));

  return window;
}

void Application::on_activate() {
  try {
    auto window = create_window();
    window->show_all();
  } catch (Glib::Error const& err) {
    std::cerr << "Application::on_activate(): " << err.what() << std::endl;
  } catch (std::exception const& err) {
    std::cerr << "Application::on_activate(): " << err.what() << std::endl;
  }
}

void Application::on_startup() { Gtk::Application::on_startup(); }
void Application::on_hide_window(Gtk::Window* window) { delete window; }