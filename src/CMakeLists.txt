file(GLOB_RECURSE SourceFile_Src "${CMAKE_CURRENT_SOURCE_DIR}/*.cc" "${CMAKE_CURRENT_SOURCE_DIR}/*.hh")
add_executable(${PROJECT_NAME} ${SourceFile_Src} ${CMAKE_BINARY_DIR}/resource.c)
target_link_libraries(${PROJECT_NAME} GTKMM HANDY pthread)
target_include_directories(${PROJECT_NAME} PRIVATE ${CMAKE_BINARY_DIR})
compile_resources(
    ${CMAKE_BINARY_DIR}/resource.c
)

install(TARGETS ${PROJECT_NAME} DESTINATION bin)