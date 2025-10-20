// --- router/index.js
import { createRouter, createWebHistory } from 'vue-router'

// --- หน้า public ---
import Homepage from '../pages/Homepage.vue'
import Documents from '../pages/Documents.vue'
import NewsEvents from '../pages/NewsEvents.vue'
import Jobs from '../pages/Jobs.vue'

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

  // ✅ หน้า dashboard

  // นักศึกษา
  { path: '/student', component: StudentDashboard },
  { path: '/student/news', component: StudentNewsEvents },
  { path: '/student/jobs', component: StudentJobs },
  { path: '/student/jobs/history', component: StudentHistoryJobs },
  { path: '/student/jobs/:jobId/apply', name: 'StudentJobApplication', component: StudentJobApplication },
  { path: '/student/documents', component: StudentDocuments },
  { path: '/student/favorites', component: StudentFavorites },
  { path: '/student/profile', component: StudentProfile },

  // อาจารย์
  { path: '/teacher', component: TeacherDashboard },
  { path: '/teacher/news', component: TeacherNewsEvents },
  { path: '/teacher/news/create', name: 'TeacherCreateNews', component: TeacherCreateNews },
  { path: '/teacher/news/edit/:id', name: 'TeacherEditNews', component: TeacherEditNews },
  { path: '/teacher/news/manage', component: TeacherManageNews },
  { path: '/teacher/jobs', component: TeacherJobs },
  { path: '/teacher/jobs/create', component: TeacherCreateJobs },

  // สถานประกอบการ
  { path: '/company', component: CompanyDashboard },
  { path: '/company/news', component: CompanyNewsEvents },
  { path: '/company/jobs', component: CompanyJobs },
  { path: '/company/jobs/manage', component: CompanyManageJobs },
  { path: '/company/jobs/create', component: CompanyCreateJobs },
  { path: '/company/jobs/applications', component: CompanyJobApplications },
  { path: '/company/evaluation/students', component: CompanyEvalutions },
  { path: '/company/documents', component: CompanyDocuments },
  { path: '/company/favorites', component: CompanyFavorites },
  { path: '/company/profile', component: CompanyProfile },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router