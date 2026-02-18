package com.netflix.accountservice.profile;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface ProfileRepository extends JpaRepository<ProfileEntity, Long> {
    List<ProfileEntity> findByUserId(Long userId);
}
