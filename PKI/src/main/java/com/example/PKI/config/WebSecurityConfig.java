package com.example.PKI.config;

import com.example.PKI.security.RestAuthenticationEntryPoint;
import com.example.PKI.security.TokenAuthenticationFilter;
import com.example.PKI.security.TokenUtils;
import com.example.PKI.service.impl.UserDetailsService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.HttpMethod;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;
import org.springframework.security.config.annotation.method.configuration.EnableGlobalMethodSecurity;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.builders.WebSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.security.web.authentication.www.BasicAuthenticationFilter;

@Configuration
@EnableGlobalMethodSecurity(prePostEnabled = true)
@EnableWebSecurity
public class WebSecurityConfig extends WebSecurityConfigurerAdapter {

    @Autowired
    private UserDetailsService customUserDetailsService;

    @Autowired
    private RestAuthenticationEntryPoint restAuthenticationEntryPoint;

    @Autowired
    private PasswordEncoder passwordEncoder;


    @Bean
    @Override
    public AuthenticationManager authenticationManagerBean() throws Exception {
        return super.authenticationManagerBean();
    }

    @Autowired
    public void configureGlobal(AuthenticationManagerBuilder auth) throws Exception {
        auth.userDetailsService(customUserDetailsService).passwordEncoder(passwordEncoder);
    }

    @Autowired
    private TokenUtils tokenUtils;

    @Override
    protected void configure(HttpSecurity http) throws Exception {
        http
                .sessionManagement().sessionCreationPolicy(SessionCreationPolicy.STATELESS).and()
                .exceptionHandling().authenticationEntryPoint(restAuthenticationEntryPoint).and()
                .authorizeRequests().antMatchers("/api/login").permitAll()
                .antMatchers("/api/confirmAccount/*").permitAll()
                .antMatchers("/api/changePassword").permitAll()
                .antMatchers("/api/sendCode").permitAll()
                .antMatchers("/api/createSubject").permitAll()
                .antMatchers("/api/checkCode").permitAll()
                .antMatchers("/api/resetPassword").permitAll()
                .antMatchers("/h2-console/**").permitAll()	// /h2-console/** ako se koristi H2 baza)
                .antMatchers("/api/foo").permitAll()
                .antMatchers("/api/two-factor-auth-status/*").permitAll()
                .antMatchers("/api/password-less-login/*").permitAll()
                .antMatchers("/api/password-less-login").permitAll()
                .anyRequest().authenticated().and()
                .cors().and()
                .addFilterBefore(new TokenAuthenticationFilter(tokenUtils,authenticationManager(), customUserDetailsService), BasicAuthenticationFilter.class)
                .headers()
                // xss protection can detect what looks like xss and blocks it - browser wont render if it sees it
                .xssProtection()
                .and()
                .contentSecurityPolicy("script-src 'self'")
                .and();
                http.csrf().disable();
    }

    // Definisanje konfiguracije koja utice na generalnu bezbednost aplikacije
    @Override
    public void configure(WebSecurity web) throws Exception {
        // Autentifikacija ce biti ignorisana ispod navedenih putanja (kako bismo ubrzali pristup resursima)
        // Zahtevi koji se mecuju za web.ignoring().antMatchers() nemaju pristup SecurityContext-u

        // Dozvoljena POST metoda na ruti /auth/login, za svaki drugi tip HTTP metode greska je 401 Unauthorized
        web.ignoring().antMatchers(HttpMethod.POST, "/api/login");
        web.ignoring().antMatchers(HttpMethod.POST, "/api/createObject");
        web.ignoring().antMatchers(HttpMethod.GET, "/api/confirmAccount/*");


        // Ovim smo dozvolili pristup statickim resursima aplikacije
        web.ignoring().antMatchers(HttpMethod.GET, "/", "/webjars/**", "/*.html", "favicon.ico", "/**/*.html",
                "/**/*.css", "/**/*.js");
    }

}



