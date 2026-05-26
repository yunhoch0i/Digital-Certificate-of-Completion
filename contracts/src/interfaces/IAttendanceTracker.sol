// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface IAttendanceTracker {
    function attendance(uint256 memberId) external view returns (uint256);
}
