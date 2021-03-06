package com.example.PKI.model;

import javax.persistence.*;
import java.security.SecureRandom;
import java.util.*;

import lombok.Data;
import org.apache.commons.codec.binary.Base32;
import org.hibernate.annotations.Fetch;
import org.hibernate.annotations.FetchMode;

@Data
@Entity
@Table(name="users")
public class User {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(unique = true, nullable = false)
    private Integer id;
    @Column(unique = true, nullable = false)
    private String email;
    @Column(nullable = false)
    private String password;
    @Column(nullable = false)
    private Role role;
    @Column
    private boolean isActivated;
    @Column
    private String commonName;
    @Column
    private String organization;
    @Column
    private String organizationUnit;
    @Column
    private String locality;
    @Column
    private String country;
    @Column(nullable = false)
    private String recoveryEmail;
    @ManyToMany(fetch = FetchType.EAGER)
    @JoinTable(name = "user_authorities",
            joinColumns = @JoinColumn(name = "user_id", referencedColumnName = "id"),
            inverseJoinColumns = @JoinColumn(name = "authority_id", referencedColumnName = "id"))
    @Fetch(value = FetchMode.SUBSELECT)
    private List<Authority> authorities;
    private boolean isUsing2FA;
    private String secret;

    public User() {  }

    public User(String email, String password, String commonName, String organization, String organizationUnit, String locality, String country,String recoveryEmail) {
        this.email = email;
        this.password = password;
        this.role = Role.UserOrdinary;
        this.isActivated = false;
        this.commonName = commonName;
        this.organization = organization;
        this.organizationUnit = organizationUnit;
        this.locality = locality;
        this.country = country;
        this.recoveryEmail = recoveryEmail;
        this.isUsing2FA = false;
        this.secret = generateSecretKey();
    }

    private static String generateSecretKey() {
        SecureRandom random = new SecureRandom();
        byte[] bytes = new byte[20];
        random.nextBytes(bytes);
        Base32 base32 = new Base32();
        return base32.encodeToString(bytes);
    }

    public void setSecret() {
        this.secret = generateSecretKey();
    }



    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }


    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public String getCommonName() {
        return commonName;
    }

    public void setCommonName(String commonName) {
        this.commonName = commonName;
    }

    public String getOrganization() {
        return organization;
    }

    public void setOrganization(String organization) {
        this.organization = organization;
    }

    public String getOrganizationUnit() {
        return organizationUnit;
    }

    public void setOrganizationUnit(String organizationUnit) {
        this.organizationUnit = organizationUnit;
    }

    public String getLocality() {
        return locality;
    }

    public void setLocality(String locality) {
        this.locality = locality;
    }

    public String getCountry() {
        return country;
    }

    public void setCountry(String country) {
        this.country = country;
    }

    public Role getRole() {
        return role;
    }

    public void setRole(Role role) {
        this.role = role;
    }

    public boolean isActivated() {
        return isActivated;
    }

    public void setActivated(boolean activated) {
        isActivated = activated;
    }

    public String getRecoveryEmail() {
        return recoveryEmail;
    }

    public void setRecoveryEmail(String recoveryEmail) {
        this.recoveryEmail = recoveryEmail;
    }
}
