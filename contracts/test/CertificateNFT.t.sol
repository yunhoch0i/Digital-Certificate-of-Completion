// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Test.sol";
import "../src/AttendanceTracker.sol";
import "../src/CertificateNFT.sol";

contract CertificateNFTTest is Test {
    AttendanceTracker tracker;
    CertificateNFT nft;

    uint256 constant THRESHOLD = 18;
    uint256 constant PASS_ID   = 0;
    uint256 constant FAIL_ID   = 1;

    function setUp() public {
        tracker = new AttendanceTracker();
        nft = new CertificateNFT(address(tracker), THRESHOLD);

        uint256[] memory passIds = new uint256[](1);
        passIds[0] = PASS_ID;
        for (uint256 i = 0; i < THRESHOLD; i++) {
            tracker.mintAttendance(passIds, "setup");
        }

        uint256[] memory failIds = new uint256[](1);
        failIds[0] = FAIL_ID;
        for (uint256 i = 0; i < THRESHOLD - 1; i++) {
            tracker.mintAttendance(failIds, "setup");
        }
    }

    function test_IssueCertificateToQualifiedMember() public {
        nft.issueCertificate(PASS_ID, "QmTest123", unicode"홍길동");
        assertTrue(nft.hasCertificate(PASS_ID));
        assertEq(nft.ownerOf(1), address(this)); // NFT goes to operator
        assertEq(nft.tokenURI(1), "ipfs://QmTest123");
        assertEq(nft.memberTokenId(PASS_ID), 1);
    }

    function test_RevertBelowThreshold() public {
        vm.expectRevert("CERT: below completion threshold");
        nft.issueCertificate(FAIL_ID, "QmFail", unicode"김철수");
    }

    function test_RevertDuplicateIssuance() public {
        nft.issueCertificate(PASS_ID, "QmTest123", unicode"홍길동");
        vm.expectRevert("CERT: already issued");
        nft.issueCertificate(PASS_ID, "QmTest456", unicode"홍길동");
    }

    function test_EmitCertificateIssuedEvent() public {
        vm.expectEmit(true, false, false, true);
        emit CertificateNFT.CertificateIssued(PASS_ID, 1, unicode"홍길동");
        nft.issueCertificate(PASS_ID, "QmTest123", unicode"홍길동");
    }
}
