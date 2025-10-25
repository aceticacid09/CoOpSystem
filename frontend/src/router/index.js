// frontend/src/router/index.js
import { createRouter, createWebHistory } from 'vue-router'

// --- หน้า public ---
import Homepage from '../pages/Homepage.vue'
import Documents from '../pages/Documents.vue'
import NewsEvents from '../pages/NewsEvents.vue'
import Jobs from '../pages/Jobs.vue'
import Login from '../pages/Login.vue'

// --- หน้า dashboard ---
// นักศึกษา
import StudentDashboard from '../pages/Student/StudentDasboard.vue'
import StudentNewsEvents from '../pages/Student/StudentNewsEvents.vue'
import StudentJobs from '../pages/Student/StudentJobs.vue'
import StudentDocuments from '../pages/Student/StudentDocuments.vue'
import StudentFavorites from '../pages/Student/StudentFavorites.vue'
import StudentProfile from '../pages/Student/StudentProfile.vue'
import StudentHistoryJobs from '../pages/Student/StudentHistoryJobs.vue'
import StudentJobApplication from '../pages/Student/StudentJobApplications.vue'

// อาจารย์
import TeacherDashboard from '../pages/Teacher/TeacherDasboard.vue'
import TeacherNewsEvents from '../pages/Teacher/TeacherNewsEvents.vue'
import TeacherCreateNews from '../pages/Teacher/TeacherCreateNews.vue'
import TeacherEditNews from '../pages/Teacher/TeacherEditNews.vue'
import TeacherManageNews from '../pages/Teacher/TeacherManageNews.vue'
import TeacherJobs from '../pages/Teacher/TeacherJobs.vue'
import TeacherCreateJobs from '../pages/Teacher/TeacherCreateJobs.vue'

// สถานประกอบการ
import CompanyDashboard from '../pages/Company/CompanyDasboard.vue'
import CompanyNewsEvents from '../pages/Company/CompanyNewsEvents.vue'
import CompanyJobs from '../pages/Company/CompanyJobs.vue'
import CompanyCreateJobs from '../pages/Company/CompanyCreateJobs.vue'
import CompanyJobApplications from '../pages/Company/CompanyJobApplications.vue'
import CompanyEvalutions from '../pages/Company/CompanyEvalutions.vue'
import CompanyDocuments from '../pages/Company/CompanyDocuments.vue'
import CompanyFavorites from '../pages/Company/CompanyFavorites.vue'
import CompanyProfile from '../pages/Company/CompanyProfile.vue'
import CompanyManageJobs from '../pages/Company/CompanyManageJobs.vue'

const routes = [
  // ✅ หน้า public
  { path: '/', component: Homepage },
  { path: '/documents', component: Documents },
  { path: '/news-events', component: NewsEvents },
  { path: '/jobs', component: Jobs },
  { path: '/login', component: Login },

  // ✅ หน้า dashboard - เพิ่ม meta: { requiresAuth: true, role: 'student' }

  // นักศึกษา
  {
    path: '/student',
    component: StudentDashboard,
    meta: { requiresAuth: true, role: 'student' }
  },
  {
    path: '/student/news',
    component: StudentNewsEvents,
    meta: { requiresAuth: true, role: 'student' }
  },
  {
    path: '/student/jobs',
    component: StudentJobs,
    meta: { requiresAuth: true, role: 'student' }
  },
  {
    path: '/student/jobs/history',
    component: StudentHistoryJobs,
    meta: { requiresAuth: true, role: 'student' }
  },
  {
    path: '/student/jobs/:jobId/apply',
    name: 'StudentJobApplication',
    component: StudentJobApplication,
    meta: { requiresAuth: true, role: 'student' }
  },
  {
    path: '/student/documents',
    component: StudentDocuments,
    meta: { requiresAuth: true, role: 'student' }
  },
  {
    path: '/student/favorites',
    component: StudentFavorites,
    meta: { requiresAuth: true, role: 'student' }
  },
  {
    path: '/student/profile',
    component: StudentProfile,
    meta: { requiresAuth: true, role: 'student' }
  },

  // อาจารย์
  {
    path: '/teacher',
    component: TeacherDashboard,
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/teacher/news',
    name: 'TeacherNewsEvents',
    component: TeacherNewsEvents,
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/teacher/news/create',
    name: 'TeacherCreateNews',
    component: TeacherCreateNews,
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/teacher/news/edit/:id',
    name: 'TeacherEditNews',
    component: TeacherEditNews,
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/teacher/news/manage',
    name: 'TeacherManageNews',
    component: TeacherManageNews,
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/teacher/jobs',
    component: TeacherJobs,
    meta: { requiresAuth: true, role: 'teacher' }
  },
  {
    path: '/teacher/jobs/create',
    component: TeacherCreateJobs,
    meta: { requiresAuth: true, role: 'teacher' }
  },

  // สถานประกอบการ
  {
    path: '/company',
    component: CompanyDashboard,
    meta: { requiresAuth: true, role: 'company' }
  },
  {
    path: '/company/news',
    component: CompanyNewsEvents,
    meta: { requiresAuth: true, role: 'company' }
  },
  {
    path: '/company/jobs',
    component: CompanyJobs,
    meta: { requiresAuth: true, role: 'company' }
  },
  {
    path: '/company/jobs/manage',
    component: CompanyManageJobs,
    meta: { requiresAuth: true, role: 'company' }
  },
  {
    path: '/company/jobs/create',
    component: CompanyCreateJobs,
    meta: { requiresAuth: true, role: 'company' }
  },
  {
    path: '/company/jobs/applications',
    component: CompanyJobApplications,
    meta: { requiresAuth: true, role: 'company' }
  },
  {
    path: '/company/evaluation/students',
    component: CompanyEvalutions,
    meta: { requiresAuth: true, role: 'company' }
  },
  {
    path: '/company/documents',
    component: CompanyDocuments,
    meta: { requiresAuth: true, role: 'company' }
  },
  {
    path: '/company/favorites',
    component: CompanyFavorites,
    meta: { requiresAuth: true, role: 'company' }
  },
  {
    path: '/company/profile',
    component: CompanyProfile,
    meta: { requiresAuth: true, role: 'company' }
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, from, next) => {
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth);
  const requiredRole = to.meta.role;
  const currentUser = localStorage.getItem('currentUser');
  const userRole = localStorage.getItem('userRole');

  if (requiresAuth) {
    if (!currentUser) {
      // ไม่ได้ล็อกอิน -> ไปหน้า Login
      next('/login');
    } else if (requiredRole && userRole !== requiredRole) {
      // Role ไม่ตรง -> ไปหน้า Dashboard ของตัวเอง
      next(`/${userRole}`);
    } else {
      // ผ่านทุกเงื่อนไข
      next();
    }
  } else {
    next();
  }
});

export default router