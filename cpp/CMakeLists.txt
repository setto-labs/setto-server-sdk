cmake_minimum_required(VERSION 3.20)
project(setto-server-sdk VERSION 0.1.0 LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

find_package(CURL REQUIRED)
find_package(nlohmann_json 3.11 QUIET)

add_library(setto-server-sdk
    src/client.cpp
    src/errors.cpp
)

target_include_directories(setto-server-sdk
    PUBLIC
        $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
        $<INSTALL_INTERFACE:include>
)

target_link_libraries(setto-server-sdk
    PRIVATE CURL::libcurl
)

if(nlohmann_json_FOUND)
    target_link_libraries(setto-server-sdk PRIVATE nlohmann_json::nlohmann_json)
endif()
