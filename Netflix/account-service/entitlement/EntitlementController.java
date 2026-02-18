package com.netflix.accountservice.entitlement;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/entitlements")
public class EntitlementController {

    private final EntitlementService entitlementService;

    public EntitlementController(EntitlementService entitlementService) {
        this.entitlementService = entitlementService;
    }

    // GET /entitlements/playback
    @GetMapping("/playback")
    public ResponseEntity<Boolean> canPlayback(@RequestParam Long userId) {
        return ResponseEntity.ok(entitlementService.canPlayback(userId));
    }

    // GET /entitlements/download
    @GetMapping("/download")
    public ResponseEntity<Boolean> canDownload(@RequestParam Long userId) {
        return ResponseEntity.ok(entitlementService.canDownload(userId));
    }

    // GET /entitlements/concurrency
    @GetMapping("/concurrency")
    public ResponseEntity<Integer> getConcurrencyLimit(@RequestParam Long userId) {
        return ResponseEntity.ok(entitlementService.getConcurrencyLimit(userId));
    }
}
