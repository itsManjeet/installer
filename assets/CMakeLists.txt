set(GRESOURCE_FILE ${CMAKE_BINARY_DIR}/${PROJECT_NAME}.gresource.xml)
set(GSCHEMA_FILE ${CMAKE_BINARY_DIR}/${PROJECT_NAME}.gschema.xml)
set(DESKTOP_FILE ${CMAKE_BINARY_DIR}/${APP_ID}.desktop)
set(ICON_FILE ${CMAKE_BINARY_DIR}/${APP_ID}.svg)

configure_file(gresource.xml.in ${GRESOURCE_FILE})
configure_file(gschema.xml.in ${GSCHEMA_FILE})
configure_file(desktop.in ${DESKTOP_FILE})
configure_file(icons/scalable/icon.svg ${ICON_FILE} COPYONLY)

compile_schemas(${GSCHEMA_FILE})
install(FILES ${GSCHEMA_FILE} DESTINATION share/glib-2.0/schemas)
install(FILES ${DESKTOP_FILE} DESTINATION share/applications)
install(FILES ${ICON_FILE} DESTINATION share/icons/${PROJECT_NAME})