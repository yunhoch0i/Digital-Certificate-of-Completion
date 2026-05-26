// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "./interfaces/IAttendanceTracker.sol";

/// @notice Issues completion certificates (ERC-721) keyed by memberId.
///         All NFTs are minted to the operator's wallet (msg.sender).
contract CertificateNFT is ERC721URIStorage, AccessControl {
    bytes32 public constant ISSUER_ROLE = keccak256("ISSUER_ROLE");

    IAttendanceTracker public attendanceTracker;
    uint256 public completionThreshold;
    uint256 private _tokenIdCounter;

    // memberId => issued
    mapping(uint256 => bool) public hasCertificate;
    // memberId => tokenId
    mapping(uint256 => uint256) public memberTokenId;

    event CertificateIssued(
        uint256 indexed memberId,
        uint256 tokenId,
        string memberName
    );

    constructor(
        address _tracker,
        uint256 _threshold
    ) ERC721("ASBG Certificate", "ASBGCERT") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ISSUER_ROLE, msg.sender);
        attendanceTracker = IAttendanceTracker(_tracker);
        completionThreshold = _threshold;
    }

    function issueCertificate(
        uint256 memberId,
        string calldata ipfsCid,
        string calldata memberName
    ) external onlyRole(ISSUER_ROLE) {
        require(
            attendanceTracker.attendance(memberId) >= completionThreshold,
            "CERT: below completion threshold"
        );
        require(!hasCertificate[memberId], "CERT: already issued");

        uint256 tokenId = ++_tokenIdCounter;
        _mint(msg.sender, tokenId); // operator is always EOA
        _setTokenURI(tokenId, string(abi.encodePacked("ipfs://", ipfsCid)));
        hasCertificate[memberId] = true;
        memberTokenId[memberId] = tokenId;

        emit CertificateIssued(memberId, tokenId, memberName);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721URIStorage, AccessControl)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}
