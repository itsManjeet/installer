#ifndef LISTBOX_HH
#define LISTBOX_HH

#include <gtkmm.h>

template <typename T>
class ListButton : public Gtk::Box {
 private:
  T m_Data;
  ListButton(std::string const& image, std::string const& title,
             std::string const& subtitle, T data)
      : m_Data(data) {
    set_margin_top(10);
    set_margin_bottom(10);
    set_margin_start(15);
    set_margin_end(15);

    set_spacing(35);

    m_Image.set_from_icon_name(image, Gtk::ICON_SIZE_DND);

    pack_start(m_Image, false, false);

    m_Title.set_markup("<span size=\"large\" weight=\"ultrabold\">" + title +
                       "</span>");
    m_Title.set_halign(Gtk::ALIGN_START);
    m_VBox.pack_start(m_Title, true, true);

    m_SubTitle.set_markup("<span color=\"gray\" weight=\"bold\">" + subtitle +
                          "</span>");
    m_SubTitle.set_halign(Gtk::ALIGN_START);

    m_VBox.pack_start(m_SubTitle, true, true);
    m_VBox.set_orientation(Gtk::ORIENTATION_VERTICAL);
    m_VBox.set_valign(Gtk::ALIGN_CENTER);
    m_VBox.set_halign(Gtk::ALIGN_START);
    pack_start(m_VBox, true, true);
  }

 protected:
  Gtk::Image m_Image;
  Gtk::Label m_Title, m_SubTitle;
  Gtk::Box m_VBox;

 public:
  static Glib::RefPtr<ListButton> create(std::string const& image,
                                         std::string const& title,
                                         std::string const& subtitle,
                                         T const& data) {
    return Glib::RefPtr<ListButton>(
        new ListButton(image, title, subtitle, data));
  }

  T const& data() const { return m_Data; }
};

class ListBox : public Gtk::ScrolledWindow {
 private:
  Gtk::ListBox m_ListBox;
  Gtk::Frame m_Frame;

 protected:
  void header_func(Gtk::ListBoxRow* row, Gtk::ListBoxRow* after);

 public:
  ListBox();

  Gtk::ListBox* get() { return &m_ListBox; }

  void clear();
  void append(Gtk::Widget* widget);

  Gtk::Widget* get_selected();
};

#endif