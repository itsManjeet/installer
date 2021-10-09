package main

const UI = `<?xml version="1.0" encoding="UTF-8"?>
<!-- Generated with glade 3.38.2 -->
<interface>
  <requires lib="gtk+" version="3.24"/>
  <object class="GtkTextBuffer" id="statusBuffer"/>
  <object class="GtkWindow" id="mainWindow">
    <property name="can-focus">False</property>
    <property name="default-width">800</property>
    <property name="default-height">600</property>
    <child>
      <object class="GtkStack" id="stack">
        <property name="visible">True</property>
        <property name="can-focus">False</property>
        <property name="transition-duration">450</property>
        <property name="transition-type">slide-left-right</property>
        <child>
          <object class="GtkBox">
            <property name="visible">True</property>
            <property name="can-focus">False</property>
            <property name="halign">center</property>
            <property name="valign">center</property>
            <property name="margin-start">71</property>
            <property name="margin-end">71</property>
            <property name="margin-bottom">66</property>
            <property name="orientation">vertical</property>
            <child>
              <object class="GtkImage">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <property name="margin-bottom">9</property>
                <property name="icon-name">emoji-symbols-symbolic</property>
                <property name="icon_size">6</property>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <property name="label" translatable="yes">Welcome!</property>
                <attributes>
                  <attribute name="weight" value="bold"/>
                  <attribute name="scale" value="2.1000000000000001"/>
                </attributes>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <property name="margin-top">2</property>
                <property name="margin-bottom">15</property>
                <property name="label" translatable="yes">Click 'Continue' to start installation</property>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton">
                <property name="visible">True</property>
                <property name="can-focus">True</property>
                <property name="receives-default">True</property>
                <property name="halign">center</property>
                <property name="valign">center</property>
                <signal name="clicked" handler="onContinueBtnClicked" swapped="no"/>
                <child>
                  <object class="GtkImage">
                    <property name="visible">True</property>
                    <property name="can-focus">False</property>
                    <property name="stock">gtk-go-forward</property>
                  </object>
                </child>
                <style>
                  <class name="circular"/>
                </style>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">3</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="name">welcomePage</property>
            <property name="title" translatable="yes">Welcome Page</property>
          </packing>
        </child>
        <child>
          <object class="GtkBox">
            <property name="visible">True</property>
            <property name="can-focus">False</property>
            <property name="orientation">vertical</property>
            <child>
              <object class="GtkImage">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <property name="margin-top">17</property>
                <property name="margin-bottom">12</property>
                <property name="icon-name">download-symbolic</property>
                <property name="icon_size">6</property>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <property name="label" translatable="yes">Installation</property>
                <attributes>
                  <attribute name="weight" value="bold"/>
                  <attribute name="scale" value="1.8"/>
                </attributes>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <property name="margin-bottom">20</property>
                <property name="label" translatable="yes">rlxos is now setting up into your system</property>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkTextView">
                <property name="visible">True</property>
                <property name="can-focus">True</property>
                <property name="editable">False</property>
                <property name="wrap-mode">word</property>
                <property name="left-margin">17</property>
                <property name="right-margin">17</property>
                <property name="top-margin">15</property>
                <property name="bottom-margin">15</property>
                <property name="cursor-visible">False</property>
                <property name="buffer">statusBuffer</property>
                <property name="accepts-tab">False</property>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">3</property>
              </packing>
            </child>
            <child>
              <object class="GtkProgressBar" id="progressBar">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <property name="show-text">True</property>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">4</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="name">installPage</property>
            <property name="title" translatable="yes">Install Page</property>
            <property name="position">1</property>
          </packing>
        </child>
        <child>
          <object class="GtkBox">
            <property name="visible">True</property>
            <property name="can-focus">False</property>
            <property name="halign">center</property>
            <property name="valign">center</property>
            <property name="margin-left">71</property>
            <property name="margin-right">71</property>
            <property name="margin-start">71</property>
            <property name="margin-end">71</property>
            <property name="margin-bottom">66</property>
            <property name="orientation">vertical</property>
            <child>
              <object class="GtkImage">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <property name="margin-bottom">9</property>
                <property name="icon-name">emblem-checked</property>
                <property name="icon_size">6</property>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <property name="label" translatable="yes">Success!</property>
                <attributes>
                  <attribute name="weight" value="bold"/>
                  <attribute name="scale" value="2.1000000000000001"/>
                </attributes>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can-focus">False</property>
                <property name="margin-top">2</property>
                <property name="margin-bottom">15</property>
                <property name="label" translatable="yes">rlxos is now successfully configured, 
you can now reboot into your freshly configured system</property>
                <property name="justify">center</property>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">2</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton">
                <property name="label" translatable="yes">reboot</property>
                <property name="visible">True</property>
                <property name="can-focus">True</property>
                <property name="receives-default">True</property>
                <property name="halign">center</property>
                <property name="valign">center</property>
                <signal name="clicked" handler="onRebootBtnClicked" swapped="no"/>
                <style>
                  <class name="suggested-action"/>
                  <class name="circular"/>
                </style>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">3</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="name">successPage</property>
            <property name="title" translatable="yes">Success Page</property>
            <property name="position">2</property>
          </packing>
        </child>
      </object>
    </child>
    <child type="titlebar">
      <object class="GtkHeaderBar">
        <property name="visible">True</property>
        <property name="can-focus">False</property>
        <property name="title" translatable="yes">System Installer</property>
      </object>
    </child>
  </object>
</interface>`
