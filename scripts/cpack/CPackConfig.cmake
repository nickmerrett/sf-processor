set(CPACK_PACKAGE_NAME "sfprocessor")
set(CPACK_PACKAGE_CONTACT "sysflow.io")
set(CPACK_PACKAGE_VENDOR "SysFlow")
set(CPACK_PACKAGE_DESCRIPTION "The SysFlow Processor implements a pluggable stream-processing pipeline and contains a built-in policy engine that evaluates rules on the ingested SysFlow stream")
set(CPACK_PACKAGE_DESCRIPTION_SUMMARY "SysFlow stream processing agent")
set(CPACK_RESOURCE_FILE_LICENSE "${CMAKE_CURRENT_LIST_DIR}/build/LICENSE.md")
set(CPACK_RESOURCE_FILE_README "${CMAKE_CURRENT_LIST_DIR}/build/README.md")
set(CPACK_STRIP_FILES "ON")
set(CPACK_PACKAGE_RELOCATABLE "OFF")

set(CPACK_PACKAGE_VERSION "$ENV{SYSFLOW_VERSION}")
if(NOT CPACK_PACKAGE_VERSION)
    set(CPACK_PACKAGE_VERSION "0.0.0")
else()
    # Remove the starting "v" in case there is one
    string(REGEX REPLACE "^v(.*)" "\\1" CPACK_PACKAGE_VERSION "${CPACK_PACKAGE_VERSION}")

    # Remove any release suffixes in case there is one
    string(REGEX REPLACE "-.*" "" CPACK_PACKAGE_VERSION "${CPACK_PACKAGE_VERSION}")
endif()
# Parse version into its major, minor, patch components
string(REGEX MATCH "^(0|[1-9][0-9]*)" CPACK_PACKAGE_VERSION_MAJOR "${CPACK_PACKAGE_VERSION}")
string(REGEX REPLACE "^(0|[1-9][0-9]*)\\.(0|[1-9][0-9]*)\\..*" "\\2" CPACK_PACKAGE_VERSION_MINOR "${CPACK_PACKAGE_VERSION}")
string(REGEX REPLACE "^(0|[1-9][0-9]*)\\.(0|[1-9][0-9]*)\\.(0|[1-9][0-9]*).*" "\\3" CPACK_PACKAGE_VERSION_PATCH "${CPACK_PACKAGE_VERSION}")

if(NOT CPACK_GENERATOR)
    set(CPACK_GENERATOR DEB RPM TGZ)
endif()

message(STATUS "Packaging SysFlow ${CPACK_PACKAGE_VERSION}")
message(STATUS "Using package generators: ${CPACK_GENERATOR}")
message(STATUS "Package architecture: ${CMAKE_SYSTEM_PROCESSOR}")

# DEB
set(CPACK_DEBIAN_PACKAGE_SECTION "utils")
if(${CMAKE_SYSTEM_PROCESSOR} STREQUAL "x86_64")
    set(CPACK_DEBIAN_PACKAGE_ARCHITECTURE "amd64")
endif()
if(${CMAKE_SYSTEM_PROCESSOR} STREQUAL "aarch64")
    set(CPACK_DEBIAN_PACKAGE_ARCHITECTURE "arm64")
endif()
set(CPACK_DEBIAN_PACKAGE_HOMEPAGE "https://github.com/sysflow-telemetry/sf-processor")
set(CPACK_DEBIAN_PACKAGE_DEPENDS "dkms (>= 2.1.0.0)")

# RPM
set(CPACK_RPM_PACKAGE_LICENSE "Apache v2.0")
set(CPACK_RPM_PACKAGE_ARCHITECTURE, "${CMAKE_SYSTEM_PROCESSOR}")
set(CPACK_RPM_PACKAGE_URL "https://github.com/sysflow-telemetry/sf-processor")
set(CPACK_RPM_PACKAGE_VERSION "${CPACK_PACKAGE_VERSION}")
set(CPACK_RPM_EXCLUDE_FROM_AUTO_FILELIST_ADDITION
    /usr/src
    /usr/share/man
    /usr/share/man/man8
    /etc
    /etc/sysflow
    /usr
    /usr/bin
    /usr/share
    /usr/lib
    /usr/lib/systemd
    /usr/lib/systemd/system)
set(CPACK_RPM_PACKAGE_RELOCATABLE "OFF")

# Contents
set(CPACK_PACKAGE_FILE_NAME ${CPACK_PACKAGE_NAME}-${CPACK_PACKAGE_VERSION}-${CMAKE_SYSTEM_PROCESSOR})
set(CPACK_INSTALLED_DIRECTORIES "${CMAKE_CURRENT_LIST_DIR}/build/bin" "/usr/bin" "${CMAKE_CURRENT_LIST_DIR}/build/resources" "/etc/sysflow" "${CMAKE_CURRENT_LIST_DIR}/build/service" "/usr/lib/systemd/system")

