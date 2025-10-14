<!-- frontend/src/components/Header.vue -->
<template>
  <header class="header">
    <div class="header-container">
      <!-- =================== Logo =================== -->
      <router-link to="/" class="logo">
        <img src="../assets/SC-SU-Logo-ENG.png" alt="logo" class="logo-img" />
        <div class="logo-text">
          <h1>Science Cooperative Education System</h1>
          <p>ระบบบริหารจัดการสหกิจศึกษา คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร</p>
        </div>
      </router-link>

      <!-- =================== Dashboard Mode =================== -->
      <div v-if="isDashboard" class="dashboard-right">
        <!-- ===== Notification Bell ===== -->
        <div class="notification" @click="toggleNotificationDropdown">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            class="bell-icon"
          >
            <path
              fill="#FFD450"
              d="M18 8a6 6 0 0 0-12 0c0 7-3 9-3 9h18s-3-2-3-9z"
            />
            <path
              stroke="#B8860B"
              stroke-width="1.8"
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M13.73 21a2 2 0 0 1-3.46 0"
            />
            <circle cx="12" cy="4" r="1.2" fill="#B8860B" />
          </svg>
          <span v-if="notificationCount > 0" class="badge">{{ notificationCount }}</span>

          <!-- Dropdown -->
          <transition name="fade">
            <div v-if="isNotificationOpen" class="notification-dropdown">
              <div class="notification-header">
                <h3>การแจ้งเตือน</h3>
                <span class="notification-count">{{ notificationCount }} รายการใหม่</span>
              </div>
              <div class="notification-list">
                <div
                  v-for="notification in notifications"
                  :key="notification.id"
                  class="notification-item"
                  :class="{ unread: !notification.read }"
                >
                  <div class="notification-icon">
                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                      <circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="2" fill="none" />
                      <path d="M12 16v-4" stroke="currentColor" stroke-width="2" stroke-linecap="round" />
                      <path d="M12 8h.01" stroke="currentColor" stroke-width="2" stroke-linecap="round" />
                    </svg>
                  </div>
                  <div class="notification-content">
                    <p class="notification-title">{{ notification.title }}</p>
                    <p class="notification-time">{{ notification.time }}</p>
                  </div>
                </div>
              </div>
              <div class="notification-footer">
                <a href="#" class="view-all">ดูทั้งหมด</a>
              </div>
            </div>
          </transition>
        </div>

        <!-- ===== Profile Dropdown ===== -->
        <div class="profile-dropdown" @click="toggleProfileDropdown">
          <div class="profile-info">
            <div class="avatar">
              <img :src="userAvatar" alt="profile" />
            </div>
            <div class="user-details">
              <p class="user-name">{{ userName }}</p>
              <p class="user-role">{{ userRoleText }}</p>
            </div>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              class="chevron-icon"
              :class="{ open: isProfileOpen }"
            >
              <polyline points="6 9 12 15 18 9"></polyline>
            </svg>
          </div>

          <!-- Profile Menu -->
          <transition name="fade">
            <div v-if="isProfileOpen" class="profile-menu">
              <router-link
                :to="`/${currentRole}/profile`"
                class="profile-menu-item"
                @click="closeProfileDropdown"
              >
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                  <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                  <circle cx="12" cy="7" r="4"></circle>
                </svg>
                <span>ข้อมูลส่วนตัว</span>
              </router-link>

              <router-link
                :to="`/${currentRole}/settings`"
                class="profile-menu-item"
                @click="closeProfileDropdown"
              >
                <!-- ✅ เปลี่ยน SVG เป็นฟันเฟือง -->
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <circle cx="12" cy="12" r="3" />
                  <path
                    d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 1 1-2.83 2.83l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-4 0v-.09a1.65 1.65 0 0 0-1-1.51 1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 1 1-2.83-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1 0-4h.09a1.65 1.65 0 0 0 1.51-1 1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 1 1 2.83-2.83l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 4 0v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 1 1 2.83 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 0 4h-.09a1.65 1.65 0 0 0-1.51 1z"
                  />
                </svg>
                <span>ตั้งค่า</span>
              </router-link>

              <div class="divider"></div>
              <a href="#" class="profile-menu-item logout" @click.prevent="handleLogout">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                  <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"></path>
                  <polyline points="16 17 21 12 16 7"></polyline>
                  <line x1="21" y1="12" x2="9" y2="12"></line>
                </svg>
                <span>ออกจากระบบ</span>
              </a>
            </div>
          </transition>
        </div>
      </div>

      <!-- =================== Normal Mode =================== -->
      <template v-else>
        <button
          class="hamburger"
          @click="toggleMenu"
          :aria-expanded="isOpen.toString()"
          aria-label="Toggle navigation menu"
        >
          <span class="bar"></span>
          <span class="bar"></span>
          <span class="bar"></span>
        </button>

        <nav class="nav" :class="{ open: isOpen }">
          <router-link to="/" class="nav-item" @click="closeMenu">หน้าแรก</router-link>
          <router-link to="/news-events" class="nav-item" @click="closeMenu">ข่าวสารและกิจกรรม</router-link>
          <router-link to="/documents" class="nav-item" @click="closeMenu">เอกสาร</router-link>
          <router-link to="/jobs" class="nav-item" @click="closeMenu">ค้นหางาน</router-link>

          <div class="login-dropdown" @click.stop="toggleDropdown">
            <button class="btn-login" :aria-expanded="isDropdownOpen.toString()">เข้าสู่ระบบ</button>
            <transition name="fade">
              <div v-if="isDropdownOpen" class="dropdown-menu">
                <router-link to="/student" class="dropdown-item" @click="handleLogin">นักศึกษา</router-link>
                <router-link to="/teacher" class="dropdown-item" @click="handleLogin">อาจารย์</router-link>
                <router-link to="/company" class="dropdown-item" @click="handleLogin">สถานประกอบการ</router-link>
              </div>
            </transition>
          </div>
        </nav>
      </template>
    </div>
  </header>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();

// Navigation states
const isOpen = ref(false);
const isDropdownOpen = ref(false);

// Dashboard states
const isNotificationOpen = ref(false);
const isProfileOpen = ref(false);

// Check if current route is dashboard
const isDashboard = computed(() => {
  const path = route.path;
  return path.startsWith('/student') || path.startsWith('/teacher') || path.startsWith('/company');
});

// Get current role from route
const currentRole = computed(() => {
  const path = route.path;
  if (path.startsWith('/student')) return 'student';
  if (path.startsWith('/teacher')) return 'teacher';
  if (path.startsWith('/company')) return 'company';
  return '';
});

// Mock user data based on role
const userData = {
  student: {
    name: 'อัครวิทย์ จันทรัง',
    role: 'นักศึกษา',
    avatar: 'https://i.pinimg.com/736x/4e/38/e7/4e38e73208c8a9c2410e4f1d9cb90ee5.jpg'
  },
  teacher: {
    name: 'ผศ.ดร.สัจอาภรณ์ ไววรรยา',
    role: 'อาจารย์',
    avatar: 'https://www.cp.su.ac.th/image/crop/856'
  },
  company: {
    name: 'บริษัท เทคโนโลยี จำกัด',
    role: 'สถานประกอบการ',
    avatar: 'https://i.pinimg.com/736x/97/0b/4f/970b4f30501bfe2bbc06c08bee62accf.jpg'
  }
};

const userName = computed(() => userData[currentRole.value]?.name || 'ผู้ใช้งาน');
const userRoleText = computed(() => userData[currentRole.value]?.role || '');
const userAvatar = computed(() => userData[currentRole.value]?.avatar || '');

// Mock notifications
const notificationCount = ref(6);
const notifications = ref([
  { id: 1, title: 'มีประกาศงานใหม่ที่น่าสนใจสำหรับคุณ', time: '5 นาทีที่แล้ว', read: false },
  { id: 2, title: 'อาจารย์ได้ประเมินผลการทำงานของคุณแล้ว', time: '1 ชั่วโมงที่แล้ว', read: false },
  { id: 3, title: 'กำหนดส่งรายงานในวันที่ 15 ต.ค. 2567', time: '3 ชั่วโมงที่แล้ว', read: false },
  { id: 4, title: 'มีข่าวสารใหม่จากคณะวิทยาศาสตร์', time: '1 วันที่แล้ว', read: true },
  { id: 5, title: 'คุณได้รับการอนุมัติเอกสารแล้ว', time: '2 วันที่แล้ว', read: true },
  { id: 6, title: 'อย่าลืมเข้าร่วมกิจกรรมในวันพรุ่งนี้', time: '3 วันที่แล้ว', read: true }
]);

// Navigation functions
const toggleMenu = () => {
  isOpen.value = !isOpen.value;
};

const closeMenu = () => {
  isOpen.value = false;
};

const toggleDropdown = () => {
  isDropdownOpen.value = !isDropdownOpen.value;
};

const handleLogin = () => {
  closeMenu();
  isDropdownOpen.value = false;
};

// Dashboard functions
const toggleNotificationDropdown = () => {
  isNotificationOpen.value = !isNotificationOpen.value;
  if (isNotificationOpen.value) {
    isProfileOpen.value = false;
  }
};

const toggleProfileDropdown = () => {
  isProfileOpen.value = !isProfileOpen.value;
  if (isProfileOpen.value) {
    isNotificationOpen.value = false;
  }
};

const closeProfileDropdown = () => {
  isProfileOpen.value = false;
};

const handleLogout = () => {
  // Handle logout logic here
  console.log('Logging out...');
  window.location.href = '/';
};

// Close dropdowns when clicking outside
onMounted(() => {
  document.addEventListener("click", (e) => {
    if (!e.target.closest(".login-dropdown")) {
      isDropdownOpen.value = false;
    }
    if (!e.target.closest(".notification")) {
      isNotificationOpen.value = false;
    }
    if (!e.target.closest(".profile-dropdown")) {
      isProfileOpen.value = false;
    }
  });
});
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Kanit:wght@400;500;600&family=Poppins:wght@400;600&display=swap');

* {
  font-family: 'Kanit', 'Prompt', 'Sarabun', 'Inter', 'Roboto', sans-serif;
  box-sizing: border-box;
}

.header {
  width: 100%;
  position: sticky;
  top: 0;
  z-index: 100;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.06);
  border-bottom: 6px solid #FFD450;
}

.header-container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 15px 30px;
  max-width: 100%;
}

/* ===== LOGO ===== */
.logo {
  display: flex;
  align-items: center;
  gap: 12px;
  text-decoration: none;
  margin-left: 30px;
}

.logo-img {
  width: 55px;
  height: 55px;
  object-fit: contain;
}

.logo-text h1 {
  font-size: 18px;
  font-weight: 600;
  color: #037266;
  margin: 0;
  letter-spacing: 0.3px;
}

.logo-text p {
  font-size: 13px;
  color: #555;
  margin: -3px 0 0 0;
}

/* ===== DASHBOARD RIGHT SECTION ===== */
.dashboard-right {
  display: flex;
  align-items: center;
  gap: 20px;
  margin-right: 20px;
}

/* ===== NOTIFICATION ===== */
.notification {
  position: relative;
  cursor: pointer;
  padding: 8px;
  border-radius: 50%;
  transition: background 0.2s ease;
}

.notification:hover {
  background: #f5f5f5;
}

.bell-icon {
  width: 37px;
  height: 37px;
}

.badge {
  position: absolute;
  top: 4px;
  right: 4px;
  background: #ff4444;
  color: white;
  font-size: 10px;
  font-weight: 600;
  padding: 2px 6px;
  border-radius: 10px;
  min-width: 18px;
  text-align: center;
}

/* Notification Dropdown */
.notification-dropdown {
  position: absolute;
  right: -10px;
  top: 45px;
  width: 360px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.12);
  border: 1px solid #eee;
  overflow: hidden;
}

.notification-header {
  padding: 16px 20px;
  border-bottom: 1px solid #f0f0f0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #fafafa;
}

.notification-header h3 {
  margin: 0;
  font-size: 16px;
  font-weight: 600;
  color: #333;
}

.notification-count {
  font-size: 12px;
  color: #037266;
  font-weight: 500;
}

.notification-list {
  max-height: 400px;
  overflow-y: auto;
}

.notification-item {
  display: flex;
  gap: 12px;
  padding: 14px 20px;
  border-bottom: 1px solid #f5f5f5;
  transition: background 0.2s ease;
  cursor: pointer;
}

.notification-item:hover {
  background: #f9f9f9;
}

.notification-item.unread {
  background: #f0f8f7;
}

.notification-icon {
  flex-shrink: 0;
}

.notification-icon svg {
  width: 20px;
  height: 20px;
  color: #037266;
}

.notification-content {
  flex: 1;
}

.notification-title {
  margin: 0;
  font-size: 13px;
  color: #333;
  line-height: 1.4;
}

.notification-time {
  margin: 4px 0 0 0;
  font-size: 11px;
  color: #999;
}

.notification-footer {
  padding: 12px 20px;
  text-align: center;
  border-top: 1px solid #f0f0f0;
  background: #fafafa;
}

.view-all {
  color: #037266;
  text-decoration: none;
  font-size: 13px;
  font-weight: 500;
  transition: color 0.2s ease;
}

.view-all:hover {
  color: #025952;
}

/* ===== PROFILE DROPDOWN ===== */
.profile-dropdown {
  position: relative;
  cursor: pointer;
}

.profile-info {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 8px 12px;
  border-radius: 25px;
  transition: background 0.2s ease;
}

.profile-info:hover {
  background: #f5f5f5;
}

.avatar {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  overflow: hidden;
  border: 2px solid #dcdcdc;
}

.avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.user-details {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.user-name {
  margin: 0;
  font-size: 14px;
  font-weight: 600;
  color: #333;
}

.user-role {
  margin: 0;
  font-size: 12px;
  color: #666;
}

.chevron-icon {
  width: 16px;
  height: 16px;
  color: #666;
  transition: transform 0.3s ease;
}

.chevron-icon.open {
  transform: rotate(180deg);
}

/* Profile Menu Dropdown */
.profile-menu {
  position: absolute;
  right: 0;
  top: 60px;
  width: 220px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.12);
  border: 1px solid #eee;
  overflow: hidden;
  padding: 8px 0;
}

.profile-menu-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 20px;
  text-decoration: none;
  color: #333;
  transition: background 0.2s ease;
  font-size: 14px;
}

.profile-menu-item svg {
  width: 18px;
  height: 18px;
  color: #666;
}

.profile-menu-item:hover {
  background: #f5f5f5;
}

.profile-menu-item.logout {
  color: #ff4444;
}

.profile-menu-item.logout svg {
  color: #ff4444;
}

.profile-menu-item.logout:hover {
  background: #fff5f5;
}

.divider {
  height: 1px;
  background: #f0f0f0;
  margin: 8px 0;
}

/* ===== NORMAL NAV ===== */
.nav {
  display: flex;
  align-items: center;
  gap: 25px;
  transition: all 0.3s ease;
}

.nav-item {
  text-decoration: none;
  color: #037266;
  font-weight: 500;
  transition: color 0.2s ease;
}

.nav-item:hover {
  color: #FFD450;
}

/* ===== LOGIN DROPDOWN ===== */
.login-dropdown {
  position: relative;
}

.btn-login {
  background: #FFD450;
  color: #333;
  font-weight: 600;
  font-size: 14px;
  text-decoration: none;
  padding: 8px 16px;
  border-radius: 20px;
  border: none;
  cursor: pointer;
  transition: all 0.25s ease;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.08);
}

.btn-login:hover {
  transform: scale(1.05);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.dropdown-menu {
  position: absolute;
  right: 0;
  top: 45px;
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.08);
  min-width: 160px;
  overflow: hidden;
  border: 1px solid #eee;
}

.dropdown-item {
  display: block;
  padding: 10px 16px;
  color: #037266;
  text-decoration: none;
  font-weight: 500;
  transition: background 0.2s ease, color 0.2s ease;
}

.dropdown-item:hover {
  background: #037266;
  color: #fff;
}

/* ===== FADE ANIMATION ===== */
.fade-enter-active,
.fade-leave-active {
  transition: all 0.25s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(-6px);
}

/* ===== HAMBURGER BUTTON ===== */
.hamburger {
  display: none;
  position: relative;
  width: 26px;
  height: 22px;
  background: none;
  border: none;
  cursor: pointer;
  z-index: 200;
}

.bar {
  display: block;
  width: 100%;
  height: 3px;
  background-color: #037266;
  border-radius: 2px;
  transition: all 0.3s ease;
  position: relative;
}

.bar+.bar {
  margin-top: 5px;
}

.hamburger[aria-expanded="true"] .bar:nth-child(1) {
  transform: translateY(8px) rotate(45deg);
}

.hamburger[aria-expanded="true"] .bar:nth-child(2) {
  opacity: 0;
}

.hamburger[aria-expanded="true"] .bar:nth-child(3) {
  transform: translateY(-8px) rotate(-45deg);
}

/* ===== RESPONSIVE ===== */
@media (max-width: 1200px) {
  .hamburger {
    display: block;
  }

  .nav {
    position: absolute;
    top: 70px;
    right: 20px;
    flex-direction: column;
    background: #fff;
    border: 1px solid #eee;
    border-radius: 8px;
    padding: 15px;
    display: none;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  }

  .nav.open {
    display: flex;
  }

  .nav-item {
    padding: 6px 0;
  }

  .dropdown-menu {
    position: static;
    box-shadow: none;
    border: none;
    background: #fafafa;
    border-radius: 6px;
    margin-top: 6px;
  }

  .dropdown-item {
    padding: 8px;
  }
}

@media (max-width: 768px) {
  .logo {
    margin-left: 0;
  }

  .dashboard-right {
    gap: 10px;
    margin-right: 0;
  }

  .user-details {
    display: none;
  }

  .notification-dropdown {
    width: 320px;
    right: -50px;
  }

  .profile-menu {
    right: -10px;
  }
}

@media (max-width: 500px) {
  .logo-text h1 {
    font-size: 15px;
  }

  .logo-text p {
    display: none;
  }

  .notification-dropdown {
    width: 280px;
    right: -80px;
  }

  .header-container {
    padding: 12px 15px;
  }
}
</style>