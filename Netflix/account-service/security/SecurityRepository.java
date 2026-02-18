package com.netflix.accountservice.security;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface SecurityRepository extends JpaRepository<SecurityEntity, String> {
    List<SecurityEntity> findByUserId(Long userId);
}
