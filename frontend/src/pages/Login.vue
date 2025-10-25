<!-- frontend/src/pages/Login.vue -->
<template>
  <div class="login-page">
    <Header />
    
    <div class="login-container">
      <div class="login-box">
        <!-- Logo Section -->
        <div class="login-header">
          <!-- <img src="../assets/SC-SU-Logo-ENG.png" alt="logo" class="login-logo" /> -->
          <h1 class="login-title">เข้าสู่ระบบ</h1>
          <p class="login-subtitle">ระบบบริหารจัดการสหกิจศึกษา</p>
        </div>

        <!-- Role Selection -->
        <div class="role-tabs">
          <button 
            v-for="role in roles" 
            :key="role.value"
            class="role-tab"
            :class="{ active: selectedRole === role.value }"
            @click="selectedRole = role.value"
          >
            <span class="role-icon" v-html="role.icon"></span>
            <span class="role-name">{{ role.label }}</span>
          </button>
        </div>

        <!-- Login Form -->
        <form @submit.prevent="handleLogin" class="login-form">
          <div class="form-group">
            <label for="username" class="form-label">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                <circle cx="12" cy="7" r="4"></circle>
              </svg>
              ชื่อผู้ใช้
            </label>
            <input
              type="text"
              id="username"
              v-model="credentials.username"
              placeholder="กรอกชื่อผู้ใช้"
              class="form-input"
              required
            />
          </div>

          <div class="form-group">
            <label for="password" class="form-label">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <rect x="3" y="11" width="18" height="11" rx="2" ry="2"></rect>
                <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
              </svg>
              รหัสผ่าน
            </label>
            <div class="password-wrapper">
              <input
                :type="showPassword ? 'text' : 'password'"
                id="password"
                v-model="credentials.password"
                placeholder="กรอกรหัสผ่าน"
                class="form-input"
                required
              />
              <button 
                type="button" 
                class="toggle-password"
                @click="showPassword = !showPassword"
              >
                <svg v-if="!showPassword" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path>
                  <circle cx="12" cy="12" r="3"></circle>
                </svg>
                <svg v-else xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M17.94 17.94A10.07 10.07 0 0 1 12 20c-7 0-11-8-11-8a18.45 18.45 0 0 1 5.06-5.94M9.9 4.24A9.12 9.12 0 0 1 12 4c7 0 11 8 11 8a18.5 18.5 0 0 1-2.16 3.19m-6.72-1.07a3 3 0 1 1-4.24-4.24"></path>
                  <line x1="1" y1="1" x2="23" y2="23"></line>
                </svg>
              </button>
            </div>
          </div>

          <div class="form-options">
            <label class="remember-me">
              <input type="checkbox" v-model="rememberMe" />
              <span>จดจำฉันไว้</span>
            </label>
            <a href="#" class="forgot-password">ลืมรหัสผ่าน?</a>
          </div>

          <!-- Error Message -->
          <transition name="fade">
            <div v-if="errorMessage" class="error-message">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <circle cx="12" cy="12" r="10"></circle>
                <line x1="12" y1="8" x2="12" y2="12"></line>
                <line x1="12" y1="16" x2="12.01" y2="16"></line>
              </svg>
              {{ errorMessage }}
            </div>
          </transition>

          <!-- Submit Button -->
          <button type="submit" class="btn-submit" :disabled="isLoading">
            <span v-if="!isLoading">เข้าสู่ระบบ</span>
            <span v-else class="loading-spinner">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="12" y1="2" x2="12" y2="6"></line>
                <line x1="12" y1="18" x2="12" y2="22"></line>
                <line x1="4.93" y1="4.93" x2="7.76" y2="7.76"></line>
                <line x1="16.24" y1="16.24" x2="19.07" y2="19.07"></line>
                <line x1="2" y1="12" x2="6" y2="12"></line>
                <line x1="18" y1="12" x2="22" y2="12"></line>
                <line x1="4.93" y1="19.07" x2="7.76" y2="16.24"></line>
                <line x1="16.24" y1="7.76" x2="19.07" y2="4.93"></line>
              </svg>
              กำลังเข้าสู่ระบบ...
            </span>
          </button>

          <!-- Register Link -->
          <div class="register-link">
            <p>ยังไม่มีบัญชี? <a href="#">สมัครสมาชิก</a></p>
          </div>
        </form>

      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import Header from '../components/Header.vue';
import { authenticateUser } from '../data/users';

const router = useRouter();

// Role configuration
const roles = [
  {
    value: 'student',
    label: 'นักศึกษา',
    icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
      <path d="M22 10v6M2 10l10-5 10 5-10 5z"></path>
      <path d="M6 12v5c3 3 9 3 12 0v-5"></path>
    </svg>`
  },
  {
    value: 'teacher',
    label: 'อาจารย์',
    icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
      <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
      <circle cx="9" cy="7" r="4"></circle>
      <path d="M23 21v-2a4 4 0 0 0-3-3.87"></path>
      <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
    </svg>`
  },
  {
    value: 'company',
    label: 'สถานประกอบการ',
    icon: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
      <rect x="2" y="7" width="20" height="14" rx="2" ry="2"></rect>
      <path d="M16 21V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16"></path>
    </svg>`
  }
];

// Form state
const selectedRole = ref('student');
const credentials = ref({
  username: '',
  password: ''
});
const showPassword = ref(false);
const rememberMe = ref(false);
const isLoading = ref(false);
const errorMessage = ref('');

// Handle login
const handleLogin = async () => {
  errorMessage.value = '';
  isLoading.value = true;

  // Simulate API call
  setTimeout(() => {
    const user = authenticateUser(
      credentials.value.username,
      credentials.value.password,
      selectedRole.value
    );

    if (user) {
      // Store user data (in real app, use Vuex or Pinia)
      localStorage.setItem('currentUser', JSON.stringify(user));
      localStorage.setItem('userRole', selectedRole.value);
      
      // Redirect to dashboard
      router.push(`/${selectedRole.value}`);
    } else {
      errorMessage.value = 'ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง';
    }

    isLoading.value = false;
  }, 1000);
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Kanit:wght@400;500;600&display=swap');

* {
  box-sizing: border-box;
  font-family: 'Kanit', 'Prompt', 'Sarabun', sans-serif;
}

.login-page {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background: linear-gradient(135deg, #e6f2f1 0%, #fff7e6 100%);
}

.login-container {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 40px 20px;
}

.login-box {
  background: white;
  border-radius: 20px;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.1);
  padding: 40px;
  width: 100%;
  max-width: 480px;
  animation: fadeInUp 0.5s ease;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Header */
.login-header {
  text-align: center;
  margin-bottom: 30px;
}

.login-logo {
  width: 80px;
  height: 80px;
  object-fit: contain;
  margin-bottom: 15px;
}

.login-title {
  font-size: 28px;
  font-weight: 600;
  color: #037266;
  margin: 0 0 8px 0;
}

.login-subtitle {
  font-size: 14px;
  color: #666;
  margin: 0;
}

/* Role Tabs */
.role-tabs {
  display: flex;
  gap: 10px;
  margin-bottom: 30px;
  background: #f8f8f8;
  padding: 6px;
  border-radius: 12px;
}

.role-tab {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
  padding: 12px 8px;
  background: transparent;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  color: #666;
}

.role-tab:hover {
  background: #fff;
  color: #037266;
}

.role-tab.active {
  background: #037266;
  color: white;
  box-shadow: 0 2px 8px rgba(3, 114, 102, 0.3);
}

.role-icon {
  display: flex;
  align-items: center;
  justify-content: center;
}

.role-icon svg {
  width: 24px;
  height: 24px;
}

.role-name {
  font-size: 13px;
  font-weight: 500;
}

/* Form */
.login-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-label {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  font-weight: 500;
  color: #333;
}

.form-label svg {
  width: 18px;
  height: 18px;
  color: #037266;
}

.form-input {
  width: 100%;
  padding: 12px 16px;
  border: 2px solid #e8e8e8;
  border-radius: 10px;
  font-size: 14px;
  transition: all 0.3s ease;
  font-family: inherit;
}

.form-input:focus {
  outline: none;
  border-color: #037266;
  box-shadow: 0 0 0 3px rgba(3, 114, 102, 0.1);
}

.password-wrapper {
  position: relative;
}

.password-wrapper .form-input {
  padding-right: 45px;
}

.toggle-password {
  position: absolute;
  right: 12px;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  cursor: pointer;
  padding: 4px;
  color: #666;
  transition: color 0.2s ease;
}

.toggle-password:hover {
  color: #037266;
}

.toggle-password svg {
  width: 20px;
  height: 20px;
}

/* Form Options */
.form-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 13px;
}

.remember-me {
  display: flex;
  align-items: center;
  gap: 6px;
  cursor: pointer;
  color: #666;
}

.remember-me input[type="checkbox"] {
  cursor: pointer;
  width: 16px;
  height: 16px;
}

.forgot-password {
  color: #037266;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.2s ease;
}

.forgot-password:hover {
  color: #025952;
  text-decoration: underline;
}

/* Error Message */
.error-message {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 12px 16px;
  background: #fff0f0;
  border: 1px solid #ffcccc;
  border-radius: 8px;
  color: #cc0000;
  font-size: 13px;
}

.error-message svg {
  width: 20px;
  height: 20px;
  flex-shrink: 0;
}

/* Submit Button */
.btn-submit {
  width: 100%;
  padding: 14px;
  background: #037266;
  color: white;
  border: none;
  border-radius: 10px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  font-family: inherit;
  margin-top: 10px;
}

.btn-submit:hover:not(:disabled) {
  background: #025952;
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(3, 114, 102, 0.3);
}

.btn-submit:active:not(:disabled) {
  transform: translateY(0);
}

.btn-submit:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.loading-spinner {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
}

.loading-spinner svg {
  width: 20px;
  height: 20px;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

/* Register Link */
.register-link {
  text-align: center;
  margin-top: 20px;
  padding-top: 20px;
  border-top: 1px solid #f0f0f0;
}

.register-link p {
  margin: 0;
  font-size: 14px;
  color: #666;
}

.register-link a {
  color: #037266;
  text-decoration: none;
  font-weight: 600;
  transition: color 0.2s ease;
}

.register-link a:hover {
  color: #025952;
  text-decoration: underline;
}

/* Demo Info */
.demo-info {
  margin-top: 20px;
  padding-top: 20px;
  border-top: 1px solid #f0f0f0;
}

.demo-info details {
  cursor: pointer;
}

.demo-info summary {
  font-size: 13px;
  color: #037266;
  font-weight: 500;
  padding: 8px 0;
  list-style: none;
  display: flex;
  align-items: center;
  gap: 6px;
}

.demo-info summary::-webkit-details-marker {
  display: none;
}

.demo-info summary::before {
  content: '▶';
  font-size: 10px;
  transition: transform 0.3s ease;
}

.demo-info details[open] summary::before {
  transform: rotate(90deg);
}

.demo-credentials {
  margin-top: 12px;
  padding: 12px;
  background: #f8f8f8;
  border-radius: 8px;
  font-size: 12px;
}

.demo-item {
  padding: 6px 0;
  color: #666;
}

.demo-item strong {
  color: #333;
  display: inline-block;
  width: 120px;
}

/* Fade Animation */
.fade-enter-active,
.fade-leave-active {
  transition: all 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

/* Responsive */
@media (max-width: 768px) {
  .login-box {
    padding: 30px 25px;
  }

  .login-title {
    font-size: 24px;
  }

  .role-tabs {
    flex-direction: column;
  }

  .role-tab {
    flex-direction: row;
    justify-content: center;
    padding: 10px;
  }

  .role-name {
    font-size: 14px;
  }
}

@media (max-width: 480px) {
  .login-container {
    padding: 20px 15px;
  }

  .login-box {
    padding: 25px 20px;
  }

  .login-logo {
    width: 60px;
    height: 60px;
  }

  .login-title {
    font-size: 22px;
  }

  .form-options {
    flex-direction: column;
    align-items: flex-start;
    gap: 10px;
  }
}
</style>