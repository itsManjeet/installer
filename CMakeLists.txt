cmake_minimum_required(VERSION 3.10)

project(installer VERSION 1.0)

set(APP_ID "dev.rlxos.Installer")
string(REPLACE "." "/" APP_PREFIX "/${APP_ID}/")

set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} "${CMAKE_SOURCE_DIR}/cmake")
set(LOCALE_DIR ${CMAKE_SOURCE_DIR}/locale)

include(gettextUtils)
include(glibUtils)

find_package(GTKMM)
find_package(HANDY)

if (CMAKE_BUILD_TYPE STREQUAL "Debug")
    add_definitions(-DDEBUG)
    set(DATA_DIR "${CMAKE_BINARY_DIR}")
endif ()

configure_file(config.h.in ${CMAKE_BINARY_DIR}/config.h)

add_subdirectory(assets)
add_subdirectory(src)

# file(GLOB_RECURSE CC_Files "${CMAKE_SOURCE_DIR}/src/*.cc")
# configure_gettext(
#     DOMAIN ${PROJECT_NAME}
#     TARGET_NAME ${PROJECT_NAME}_locale
#     SOURCES ${CC_Files}
#     POTFILE_DESTINATION share
#     XGETTEXT_ARGS
#         "--keyword=_" "--keyword=N_" "--keyword=P_:1,2"
#         "--package-name=${PROJECT_NAME}" "--package-version=${PROJECT_VERSION}"
#         "--copyright-holder=Manjeet Singh" "--msgid-bugs-address=rlxos.dev/bugs"
#     LANGUAGES "en_US.UTF-8" "hi_IN.UTF-8"
# )
# find_package(Intl REQUIRED)

# add_dependencies(${PROJECT_NAME} ${PROJECT_NAME}_locale)