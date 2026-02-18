package com.netflix.accountservice.user;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/users")
public class UserController {

    private final UserService userService;

    public UserController(UserService userService) {
        this.userService = userService;
    }

    // POST /users/register
    @PostMapping("/register")
    public ResponseEntity<UserEntity> register(@RequestBody UserEntity user) {
        return ResponseEntity.ok(userService.register(user));
    }

    // GET /users/{id}
    @GetMapping("/{id}")
    public ResponseEntity<UserEntity> getUser(@PathVariable Long id) {
        return userService.findById(id)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }

    // PATCH /users/{id}/status
    @PatchMapping("/{id}/status")
    public ResponseEntity<UserEntity> updateStatus(@PathVariable Long id,
                                                   @RequestParam String status) {
        return ResponseEntity.ok(userService.updateStatus(id, status));
    }

    // DELETE /users/{id}
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteUser(@PathVariable Long id) {
        userService.deleteUser(id);
        return ResponseEntity.noContent().build();
    }
}
