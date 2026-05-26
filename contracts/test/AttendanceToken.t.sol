// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Test.sol";
import "../src/AttendanceToken.sol";

contract AttendanceTokenTest is Test {
    AttendanceToken token;
    address minter = address(this);
    address member = address(0x1);
    address stranger = address(0x2);

    function setUp() public {
        token = new AttendanceToken();
    }

    function test_MintByMinterRole() public {
        token.mintAttendance(member, 1, "2025-03-01");
        assertEq(token.balanceOf(member), 1);
    }

    function test_RevertMintWithoutMinterRole() public {
        vm.prank(stranger);
        vm.expectRevert();
        token.mintAttendance(member, 1, "2025-03-01");
    }

    function test_RevertTransferByNonMinter() public {
        token.mintAttendance(member, 1, "2025-03-01");
        vm.prank(member);
        vm.expectRevert("ATK: transfer disabled");
        token.transfer(stranger, 1);
    }

    function test_EmitAttendanceMintedEvent() public {
        vm.expectEmit(true, false, false, true);
        emit AttendanceToken.AttendanceMinted(member, 1, "2025-03-01");
        token.mintAttendance(member, 1, "2025-03-01");
    }

    function test_DecimalsIsZero() public view {
        assertEq(token.decimals(), 0);
    }
}
