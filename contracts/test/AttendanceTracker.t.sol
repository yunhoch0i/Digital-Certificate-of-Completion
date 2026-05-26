// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Test.sol";
import "../src/AttendanceTracker.sol";

contract AttendanceTrackerTest is Test {
    AttendanceTracker tracker;
    address stranger = address(0x99);

    function setUp() public {
        tracker = new AttendanceTracker();
    }

    function test_BatchMintByMinterRole() public {
        uint256[] memory ids = new uint256[](3);
        ids[0] = 0; ids[1] = 1; ids[2] = 5;
        tracker.mintAttendance(ids, "2025-03-01");
        assertEq(tracker.attendance(0), 1);
        assertEq(tracker.attendance(1), 1);
        assertEq(tracker.attendance(5), 1);
        assertEq(tracker.attendance(2), 0); // absent member untouched
    }

    function test_RevertMintWithoutMinterRole() public {
        uint256[] memory ids = new uint256[](1);
        ids[0] = 0;
        vm.prank(stranger);
        vm.expectRevert();
        tracker.mintAttendance(ids, "2025-03-01");
    }

    function test_AccumulatesAcrossEvents() public {
        uint256[] memory ids = new uint256[](1);
        ids[0] = 0;
        tracker.mintAttendance(ids, "event-1");
        tracker.mintAttendance(ids, "event-2");
        assertEq(tracker.attendance(0), 2);
    }

    function test_EmitAttendanceMintedEvent() public {
        uint256[] memory ids = new uint256[](1);
        ids[0] = 3;
        vm.expectEmit(true, false, false, true);
        emit AttendanceTracker.AttendanceMinted(3, "2025-03-01");
        tracker.mintAttendance(ids, "2025-03-01");
    }

    function test_EmptyBatchIsNoop() public {
        uint256[] memory ids = new uint256[](0);
        tracker.mintAttendance(ids, "2025-03-01"); // should not revert
    }
}
