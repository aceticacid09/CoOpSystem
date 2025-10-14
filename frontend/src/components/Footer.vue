<template>
  <footer class="footer">
    <!-- Top Section -->
    <div class="footer-container">
      <!-- Logo + Text -->
      <div class="footer-left">
        <img src="../assets/SC-SU-Logo-ENG.png" alt="logo" class="footer-logo" />
        <div class="footer-text">
          <h2 class="title-en">Science Cooperative Education System</h2>
          <p class="title-th">
            ระบบบริหารจัดการสหกิจศึกษา คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร
          </p>
        </div>
      </div>

      <!-- Navigation -->
      <nav class="footer-nav">
        <router-link to="/" class="footer-nav-item">หน้าแรก</router-link>
        <router-link to="/news-events" class="footer-nav-item">ข่าวสารและกิจกรรม</router-link>
        <router-link to="/documents" class="footer-nav-item">เอกสาร</router-link>
        <router-link to="/staff" class="footer-nav-item">บุคลากร</router-link>
        <router-link to="/jobs" class="footer-nav-item">ค้นหางาน</router-link>
      </nav>
    </div>

    <hr class="footer-divider" />

    <!-- Bottom Section -->
    <div class="footer-bottom">
      <div class="social-icons">
        <a href="#" class="facebook"><i class="fab fa-facebook-f"></i></a>
        <a href="#" class="youtube"><i class="fab fa-youtube"></i></a>
        <a href="#" class="tiktok"><i class="fab fa-tiktok"></i></a>
        <a href="#" class="instagram"><i class="fab fa-instagram"></i></a>
      </div>
    </div>
  </footer>
  <!-- Yellow Bar -->
    <div class="footer-bar">
      <p class="copyright">
        © 2004 - 2025 Faculty of Science, Silpakorn University. All rights reserved.
      </p>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Kanit:wght@400;600&family=Poppins:wght@400;600&display=swap');

* {
  font-family: 'Kanit','Prompt','Sarabun','Inter','Roboto',sans-serif;
}

.footer {
  background: #f9fafb;
  color: #333;
  padding: 40px 60px 0;
}

/* Top */
.footer-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 20px;
}

.footer-left {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.footer-logo {
  width: 80px;
  margin-bottom: 8px;
}

.footer-text {
  text-align: left;
}

.title-en {
  font-size: 17px;
  font-weight: 600;
  color: #037266;
  margin: 0;
}

.title-th {
  font-size: 13px;
  color: #555;
  margin: -3px 0 0 0;
}

.footer-nav {
  display: flex;
  gap: 20px;
  flex-wrap: wrap;
}

.footer-nav-item {
  text-decoration: none;
  color: #037266;
  font-weight: 500;
  transition: color 0.2s ease;
}

.footer-nav-item:hover {
  color: #ffd966;
}

/* Divider */
.footer-divider {
  border: none;
  border-top: 1px solid #e5e5e5;
  margin: 20px 0;
}

/* Bottom */
.footer-bottom {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  padding-bottom: 15px;
}

.social-icons a {
  font-size: 18px;
  margin-left: 15px;
  transition: transform 0.2s ease;
}

.social-icons a:hover {
  transform: scale(1.1);
}

/* Brand colors */
.social-icons a.facebook { color: #1877f2; }
.social-icons a.youtube { color: #ff0000; }
.social-icons a.tiktok { color: #010101; }
.social-icons a.instagram { color: #e1306c; }

/* Yellow Bar */
.footer-bar {
  background: #f7c948;
  text-align: center;
  padding: 8px 0;
  margin-top: 15px;
}

.footer-bar .copyright {
  font-size: 13px;
  color: #333;
  margin: 0;
}

/* Responsive */
@media (max-width: 900px) {
  .footer-container {
    flex-direction: column;
    align-items: flex-start;
  }

  .footer-bottom {
    justify-content: center;
  }
}
</style>
