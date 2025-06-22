import React, { useState } from 'react';
import styles from './LoginPage.module.css';

const LoginPage: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // Handle login logic here
    console.log('Login attempt:', { email, password });
  };

  const handleSocialLogin = (provider: string) => {
    // Handle social login logic here
    console.log(`Login with ${provider}`);
  };

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <div className={styles.logo}>
          <svg className={styles.logoIcon} viewBox="0 0 24 24" fill="none">
            <path
              d="M12 2L15.09 8.26L22 9L17 14L18.18 21L12 17.77L5.82 21L7 14L2 9L8.91 8.26L12 2Z"
              fill="currentColor"
            />
          </svg>
          <span className={styles.logoText}>SnowUI</span>
        </div>
        <nav className={styles.nav}>
          <a href="#" className={styles.navLink}>Product</a>
          <a href="#" className={styles.navLink}>Solutions</a>
          <a href="#" className={styles.navLink}>Resources</a>
          <a href="#" className={styles.navLink}>Download</a>
          <a href="#" className={styles.navLink}>Pricing</a>
        </nav>
        <div className={styles.headerActions}>
          <button className={styles.signUpBtn}>Sign up</button>
          <button className={styles.signInBtn}>Sign In</button>
        </div>
      </div>

      <div className={styles.loginContainer}>
        <button className={styles.closeBtn}>×</button>
        
        <div className={styles.loginContent}>
          <div className={styles.logoSection}>
            <div className={styles.loginLogo}>
              <svg className={styles.loginLogoIcon} viewBox="0 0 24 24" fill="none">
                <path
                  d="M12 2L15.09 8.26L22 9L17 14L18.18 21L12 17.77L5.82 21L7 14L2 9L8.91 8.26L12 2Z"
                  fill="currentColor"
                />
              </svg>
            </div>
          </div>

          <h2 className={styles.title}>Sign In</h2>

          <div className={styles.socialButtons}>
            <button 
              className={`${styles.socialBtn} ${styles.googleBtn}`}
              onClick={() => handleSocialLogin('Google')}
            >
              <svg className={styles.socialIcon} viewBox="0 0 24 24">
                <path fill="#4285F4" d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/>
                <path fill="#34A853" d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/>
                <path fill="#FBBC05" d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"/>
                <path fill="#EA4335" d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/>
              </svg>
            </button>
            
            <button 
              className={`${styles.socialBtn} ${styles.appleBtn}`}
              onClick={() => handleSocialLogin('Apple')}
            >
              <svg className={styles.socialIcon} viewBox="0 0 24 24">
                <path fill="#000" d="M18.71 19.5c-.83 1.24-1.71 2.45-3.05 2.47-1.34.03-1.77-.79-3.29-.79-1.53 0-2 .77-3.27.82-1.31.05-2.3-1.32-3.14-2.53C4.25 17 2.94 12.45 4.7 9.39c.87-1.52 2.43-2.48 4.12-2.51 1.28-.02 2.5.87 3.29.87.78 0 2.26-1.07 3.81-.91.65.03 2.47.26 3.64 1.98-.09.06-2.17 1.28-2.15 3.81.03 3.02 2.65 4.03 2.68 4.04-.03.07-.42 1.44-1.38 2.83M13 3.5c.73-.83 1.94-1.46 2.94-1.5.13 1.17-.34 2.35-1.04 3.19-.69.85-1.83 1.51-2.95 1.42-.15-1.15.41-2.35 1.05-3.11z"/>
              </svg>
            </button>
            
            <button 
              className={`${styles.socialBtn} ${styles.facebookBtn}`}
              onClick={() => handleSocialLogin('Facebook')}
            >
              <svg className={styles.socialIcon} viewBox="0 0 24 24">
                <path fill="white" d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/>
              </svg>
            </button>
          </div>

          <form className={styles.form} onSubmit={handleSubmit}>
            <div className={styles.inputGroup}>
              <input
                type="email"
                placeholder="Email or Phone number"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className={styles.input}
                required
              />
            </div>
            
            <div className={styles.inputGroup}>
              <input
                type="password"
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className={styles.input}
                required
              />
            </div>

            <button type="submit" className={styles.signInButton}>
              Sign In
            </button>
          </form>

          <div className={styles.footer}>
            <a href="#" className={styles.footerLink}>Sign Up</a>
            <a href="#" className={styles.footerLink}>Forgot Password</a>
            <a href="#" className={styles.footerLink}>Contact us</a>
          </div>
        </div>

        <div className={styles.copyright}>
          © 2025 SnowUI
        </div>
      </div>
    </div>
  );
};

export default LoginPage;