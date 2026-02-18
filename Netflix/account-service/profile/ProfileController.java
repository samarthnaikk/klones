package com.netflix.accountservice.profile;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/profiles")
public class ProfileController {

    private final ProfileService profileService;

    public ProfileController(ProfileService profileService) {
        this.profileService = profileService;
    }

    // POST /profiles
    @PostMapping
    public ResponseEntity<ProfileEntity> createProfile(@RequestBody ProfileEntity profile) {
        return ResponseEntity.ok(profileService.createProfile(profile));
    }

    // GET /profiles/{id}
    @GetMapping("/{id}")
    public ResponseEntity<ProfileEntity> getProfile(@PathVariable Long id) {
        return profileService.findById(id)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }

    // GET /users/{id}/profiles
    @GetMapping("/user/{userId}")
    public ResponseEntity<List<ProfileEntity>> getUserProfiles(@PathVariable Long userId) {
        return ResponseEntity.ok(profileService.findByUserId(userId));
    }

    // PATCH /profiles/{id}
    @PatchMapping("/{id}")
    public ResponseEntity<ProfileEntity> updateProfile(@PathVariable Long id,
                                                       @RequestBody ProfileEntity profile) {
        return ResponseEntity.ok(profileService.updateProfile(id, profile));
    }

    // DELETE /profiles/{id}
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteProfile(@PathVariable Long id) {
        profileService.deleteProfile(id);
        return ResponseEntity.noContent().build();
    }
}
