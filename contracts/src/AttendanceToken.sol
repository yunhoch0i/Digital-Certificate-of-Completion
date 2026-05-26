// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

contract AttendanceToken is ERC20, AccessControl {
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    event AttendanceMinted(
        address indexed member,
        uint256 amount,
        string eventId
    );

    constructor() ERC20("ASBG Attendance Token", "ASBG") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(MINTER_ROLE, msg.sender);
    }

    function decimals() public pure override returns (uint8) {
        return 0;
    }

    function mintAttendance(
        address member,
        uint256 amount,
        string calldata eventId
    ) external onlyRole(MINTER_ROLE) {
        _mint(member, amount);
        emit AttendanceMinted(member, amount, eventId);
    }

    function _update(
        address from,
        address to,
        uint256 amount
    ) internal override {
        if (from != address(0) && !hasRole(MINTER_ROLE, from)) {
            revert("ATK: transfer disabled");
        }
        super._update(from, to, amount);
    }
}
