// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/access/AccessControl.sol";

/// @notice Tracks attendance per member by numeric ID (no individual wallets needed).
contract AttendanceTracker is AccessControl {
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    // memberId (CSV row index, 0-based) => attendance count
    mapping(uint256 => uint256) public attendance;

    event AttendanceMinted(uint256 indexed memberId, string eventId);

    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(MINTER_ROLE, msg.sender);
    }

    /// @notice Batch-mint attendance for present members in one transaction.
    function mintAttendance(
        uint256[] calldata memberIds,
        string calldata eventId
    ) external onlyRole(MINTER_ROLE) {
        for (uint256 i = 0; i < memberIds.length; i++) {
            attendance[memberIds[i]]++;
            emit AttendanceMinted(memberIds[i], eventId);
        }
    }
}
