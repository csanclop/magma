// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

#pragma once

#include <folly/futures/Future.h>

#include <devmand/channels/snmp/Channel.h>

namespace devmand {
namespace channels {
namespace snmp {

using Location = std::string;
using Contact = std::string;

struct InterfaceStatus {
  int index;
  std::string status;
};

struct InterfaceName {
  int index;
  std::string name;
};

using InterfaceStatuses = std::vector<InterfaceStatus>;
using InterfaceNames = std::vector<InterfaceName>;
using InterfaceIndicies = std::vector<int>;

class IfMib {
 public:
  IfMib() = delete;
  ~IfMib() = delete;
  IfMib(const IfMib&) = delete;
  IfMib& operator=(const IfMib&) = delete;
  IfMib(IfMib&&) = delete;
  IfMib& operator=(IfMib&&) = delete;

 public:
  static folly::Future<int> getNumberOfInterfaces(
      channels::snmp::Channel& channel);
  static folly::Future<std::string> getSystemName(
      channels::snmp::Channel& channel);
  static folly::Future<Contact> getSystemContact(
      channels::snmp::Channel& channel);
  static folly::Future<Location> getSystemLocation(
      channels::snmp::Channel& channel);
  static folly::Future<InterfaceIndicies> getInterfaceIndicies(
      channels::snmp::Channel& channel);
  static folly::Future<InterfaceNames> getInterfaceNames(
      channels::snmp::Channel& channel);
  static folly::Future<InterfaceStatuses> getInterfaceStatuses(
      channels::snmp::Channel& channel);

  static folly::Future<InterfaceIndicies> handleNextInterfaceIndex(
      channels::snmp::Channel& channel,
      int numInterfacesRemaining,
      InterfaceIndicies indicies,
      channels::snmp::Oid marker);
};

} // namespace snmp
} // namespace channels
} // namespace devmand