package com.netflix.accountservice.profile;

import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class ProfileService {

    private final ProfileRepository profileRepository;

    public ProfileService(ProfileRepository profileRepository) {
        this.profileRepository = profileRepository;
    }

    public ProfileEntity createProfile(ProfileEntity profile) {
        return profileRepository.save(profile);
    }

    public Optional<ProfileEntity> findById(Long id) {
        return profileRepository.findById(id);
    }

    public List<ProfileEntity> findByUserId(Long userId) {
        return profileRepository.findByUserId(userId);
    }

    public ProfileEntity updateProfile(Long id, ProfileEntity updated) {
        ProfileEntity profile = profileRepository.findById(id)
                .orElseThrow(() -> new RuntimeException("Profile not found: " + id));
        profile.setName(updated.getName());
        profile.setAvatarUrl(updated.getAvatarUrl());
        profile.setKid(updated.isKid());
        return profileRepository.save(profile);
    }

    public void deleteProfile(Long id) {
        profileRepository.deleteById(id);
    }
}
