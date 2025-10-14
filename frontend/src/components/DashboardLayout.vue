<!-- frontend/src/components/DashboardLayout.vue -->
<template>
  <Header />
  <div class="dashboard">
    <!-- Sidebar -->
    <aside class="sidebar" :class="{ collapsed: !isSidebarOpen }">
      <nav class="menu">
        <template v-for="item in currentMenu" :key="item.name">
          <!-- Menu item without submenu -->
          <router-link v-if="!item.submenu" :to="item.link" class="menu-item" :class="{ active: isActive(item.link) }">
            <span class="icon" v-html="item.icon"></span>
            <span class="label" v-if="isSidebarOpen">{{ item.name }}</span>
          </router-link>

          <!-- Menu item with submenu -->
          <div v-else class="menu-group">
            <div class="menu-item parent" :class="{
              active: isParentActive(item),
              expanded: expandedMenus.includes(item.name)
            }" @click="toggleSubmenu(item.name)">
              <span class="icon" v-html="item.icon"></span>
              <span class="label" v-if="isSidebarOpen">{{ item.name }}</span>
              <span class="arrow" v-if="isSidebarOpen">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor"
                  stroke-width="2">
                  <polyline points="6 9 12 15 18 9"></polyline>
                </svg>
              </span>
            </div>

            <!-- Submenu -->
            <transition name="submenu">
              <div v-if="expandedMenus.includes(item.name) && isSidebarOpen" class="submenu">
                <router-link v-for="subitem in item.submenu" :key="subitem.name" :to="subitem.link" class="submenu-item"
                  :class="{ active: isActive(subitem.link) }">
                  {{ subitem.name }}
                </router-link>
              </div>
            </transition>
          </div>
        </template>
      </nav>
    </aside>

    <!-- Main content -->
    <div class="main-wrapper">
      <!-- Topbar -->
      <header class="topbar">
        <button class="hamburger" @click="toggleSidebar">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"
            stroke-linecap="round" stroke-linejoin="round" class="hamburger-icon">
            <line x1="3" y1="6" x2="21" y2="6"></line>
            <line x1="3" y1="12" x2="21" y2="12"></line>
            <line x1="3" y1="18" x2="21" y2="18"></line>
          </svg>
        </button>

        <!-- Breadcrumb Section -->
        <div class="breadcrumb-container">
          <div class="breadcrumb-line">
            <span class="breadcrumb-title">{{ currentPageTitle }}</span>
            <span class="breadcrumb-divider">|</span>
            <div class="breadcrumb-path">
              <router-link :to="homeLink" class="breadcrumb-link">หน้าแรก</router-link>
              <template v-if="breadcrumbItems.length > 0">
                <template v-for="(crumb, index) in breadcrumbItems" :key="index">
                  <span class="breadcrumb-arrow">></span>
                  <router-link v-if="crumb.link && index < breadcrumbItems.length - 1" :to="crumb.link"
                    class="breadcrumb-link">
                    {{ crumb.name }}
                  </router-link>
                  <span v-else class="breadcrumb-current">{{ crumb.name }}</span>
                </template>
              </template>
            </div>
          </div>
        </div>
      </header>

      <!-- Content Area with White Background -->
      <div class="content-wrapper">
        <main class="main">
          <slot />
        </main>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue';
import { useRoute } from 'vue-router';
import Header from '../components/Header.vue';

const props = defineProps({
  role: {
    type: String,
    default: 'student', // student, teacher, company
    validator: (value) => ['student', 'teacher', 'company'].includes(value)
  }
});

const route = useRoute();
const isSidebarOpen = ref(true);
const expandedMenus = ref([]);

// SVG Icons
const icons = {
  home: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9.75L12 3l9 6.75V21a.75.75 0 01-.75.75h-4.5a.75.75 0 01-.75-.75v-4.5H9v4.5a.75.75 0 01-.75.75h-4.5A.75.75 0 013 21V9.75z"/></svg>`,
  news: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h7"/></svg>`,
  announce: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5.882V19.24a1.76 1.76 0 01-3.417.592l-2.147-6.15M18 13a3 3 0 100-6M5.436 13.683A4.001 4.001 0 017 6h1.832c4.1 0 7.625-1.234 9.168-3v14c-1.543-1.766-5.067-3-9.168-3H7a3.988 3.988 0 01-1.564-.317z"/></svg>`,
  user: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12c2.28 0 4.13-1.85 4.13-4.13S14.28 3.75 12 3.75 7.88 5.6 7.88 7.88 9.72 12 12 12z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.5 20.25a7.5 7.5 0 0115 0"/></svg>`,
  fav: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21l-1.45-1.32C5.4 15.36 2 12.28 2 8.5A4.5 4.5 0 016.5 4c1.74 0 3.41.81 4.5 2.09A6.25 6.25 0 0115.5 4 4.5 4.5 0 0120 8.5c0 3.78-3.4 6.86-8.55 11.18L12 21z"/></svg>`,
  doc: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.8"><path stroke-linecap="round" stroke-linejoin="round" d="M6 3a2 2 0 00-2 2v14a2 2 0 002 2h12a2 2 0 002-2V9l-6-6H6z" /><path stroke-linecap="round" stroke-linejoin="round" d="M14 3v6h6" /><line x1="8" y1="13" x2="16" y2="13" stroke="currentColor" stroke-linecap="round"/><line x1="8" y1="17" x2="16" y2="17" stroke="currentColor" stroke-linecap="round"/></svg>`,
  manage: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"/></svg>`,
  evaluation: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4"/></svg>`
};

// Menus for different roles
const menus = {
  student: [
    { name: 'หน้าแรก', link: '/student', icon: icons.home },
    { name: 'ข่าวสารและกิจกรรม', link: '/student/news', icon: icons.news },
    {
      name: 'ประกาศงาน',
      icon: icons.announce,
      submenu: [
        { name: 'ค้นหางาน', link: '/student/jobs' },
        { name: 'ประวัติการสมัครงาน', link: '/student/jobs/history' }
      ]
    },
    { name: 'ข้อมูลส่วนตัว', link: '/student/profile', icon: icons.user },
    { name: 'รายการโปรด', link: '/student/favorites', icon: icons.fav },
    { name: 'เอกสาร', link: '/student/documents', icon: icons.doc }
  ],
  teacher: [
    { name: 'หน้าแรก', link: '/teacher', icon: icons.home },
    {
      name: 'ข่าวสารและกิจกรรม',
      icon: icons.news,
      submenu: [
        { name: 'ค้นหาข่าวสาร', link: '/teacher/news' },
        { name: 'สร้างข่าวสารใหม่', link: '/teacher/news/create' }
      ]
    },
    {
      name: 'ประกาศงาน',
      icon: icons.announce,
      submenu: [
        { name: 'ค้นหางาน', link: '/teacher/jobs' },
        { name: 'สร้างประกาศงานใหม่', link: '/teacher/jobs/create' }
      ]
    },
    {
      name: 'การจัดการข้อมูล',
      icon: icons.manage,
      submenu: [
        { name: 'ข้อมูลนักศึกษา', link: '/teacher/manage/students' },
        { name: 'ข้อมูลสถานประกอบการ', link: '/teacher/manage/companies' }
      ]
    },
    {
      name: 'ระบบการประเมิน',
      icon: icons.evaluation,
      submenu: [
        { name: 'ประเมินนักศึกษา', link: '/teacher/evaluation/students' },
        { name: 'ประเมินสถานประกอบการ', link: '/teacher/evaluation/companies' }
      ]
    },
    { name: 'ข้อมูลส่วนตัว', link: '/teacher/profile', icon: icons.user },
    { name: 'รายการโปรด', link: '/teacher/favorites', icon: icons.fav },
    { name: 'เอกสาร', link: '/teacher/documents', icon: icons.doc }
  ],
  company: [
    { name: 'หน้าแรก', link: '/company', icon: icons.home },
    { name: 'ข่าวสารและกิจกรรม', link: '/company/news', icon: icons.news },
    {
      name: 'ประกาศงาน',
      icon: icons.announce,
      submenu: [
        { name: 'ค้นหางาน', link: '/company/jobs' },
        { name: 'ประวัติประกาศงาน', link: '/company/jobs/history' },
        { name: 'สร้างประกาศงานใหม่', link: '/company/jobs/create' },
        { name: 'ใบสมัครที่ได้รับ', link: '/company/jobs/applications' }
      ]
    },
    {
      name: 'ระบบการประเมิน',
      icon: icons.evaluation,
      submenu: [
        { name: 'ประเมินนักศึกษา', link: '/company/evaluation/students' }
      ]
    },
    { name: 'รายการโปรด', link: '/company/favorites', icon: icons.fav },
    { name: 'เอกสาร', link: '/company/documents', icon: icons.doc }
  ]
};

const currentMenu = computed(() => menus[props.role]);
const homeLink = computed(() => `/${props.role}`);

const isActive = (path) => route.path === path;

const isParentActive = (item) => {
  if (!item.submenu) return false;
  return item.submenu.some(sub => route.path === sub.link);
};

const toggleSubmenu = (menuName) => {
  const index = expandedMenus.value.indexOf(menuName);
  if (index > -1) {
    expandedMenus.value.splice(index, 1);
  } else {
    expandedMenus.value.push(menuName);
  }
};

const toggleSidebar = () => (isSidebarOpen.value = !isSidebarOpen.value);

// Auto-expand parent menu if child is active
watch(() => route.path, () => {
  currentMenu.value.forEach(item => {
    if (item.submenu && isParentActive(item)) {
      if (!expandedMenus.value.includes(item.name)) {
        expandedMenus.value.push(item.name);
      }
    }
  });
}, { immediate: true });

// Breadcrumb logic
const breadcrumbItems = computed(() => {
  const items = [];

  for (const item of currentMenu.value) {
    if (item.link === route.path) {
      if (route.path !== homeLink.value) {
        items.push({ name: item.name, link: null });
      }
      break;
    }

    if (item.submenu) {
      for (const subitem of item.submenu) {
        if (subitem.link === route.path) {
          items.push({ name: item.name, link: null });
          items.push({ name: subitem.name, link: null });
          break;
        }
      }
    }
  }

  return items;
});

const currentPageTitle = computed(() => {
  if (breadcrumbItems.value.length > 0) {
    return breadcrumbItems.value[breadcrumbItems.value.length - 1].name;
  }

  const found = currentMenu.value.find((m) => m.link === route.path);
  return found ? found.name : 'หน้าแรก';
});
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Kanit:wght@400;700&display=swap");

* {
  box-sizing: border-box;
  font-family: "Kanit", "Prompt", "Sarabun", "Noto Sans Thai", sans-serif;
}

.dashboard {
  display: flex;
  min-height: 100vh;
  background: #f5f5f5;
  font-family: 'Kanit', 'Poppins', sans-serif;
}

/* Sidebar */
.sidebar {
  width: 250px;
  background: #ffffff;
  border-right: 1px solid #eee;
  padding-top: 10px;
  transition: width 0.3s ease;
  box-shadow: 2px 0 4px rgba(0, 0, 0, 0.02);
  overflow-y: auto;
}

.sidebar.collapsed {
  width: 70px;
}

.menu {
  display: flex;
  flex-direction: column;
}

.menu-group {
  margin: 2px 0;
}

.menu-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  margin: 4px 8px;
  border-radius: 8px;
  text-decoration: none;
  color: #666;
  transition: all 0.25s ease;
  font-size: 14px;
  cursor: pointer;
  position: relative;
}

.menu-item.parent {
  padding-right: 36px;
}

.menu-item .icon {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 22px;
  height: 22px;
  color: #999;
  transition: color 0.25s ease;
  flex-shrink: 0;
}

.menu-item .arrow {
  position: absolute;
  right: 12px;
  width: 16px;
  height: 16px;
  color: #999;
  transition: transform 0.3s ease, color 0.25s ease;
}

.menu-item.expanded .arrow {
  transform: rotate(180deg);
}

.menu-item.active {
  background: #037266;
  color: #fff;
  box-shadow: 0 2px 4px rgba(3, 114, 102, 0.2);
}

.menu-item.active .icon,
.menu-item.active .arrow {
  color: #fff;
}

.menu-item:hover {
  background: #f0f8f5;
  color: #037266;
}

.menu-item:hover .icon,
.menu-item:hover .arrow {
  color: #037266;
}

.menu-item.active:hover {
  background: #037266;
  color: #fff;
}

.menu-item.active:hover .icon,
.menu-item.active:hover .arrow {
  color: #fff;
}

.label {
  white-space: nowrap;
  font-weight: 400;
  flex: 1;
}

/* Submenu */
.submenu {
  display: flex;
  flex-direction: column;
  margin: 0 8px 8px 8px;
  padding-left: 38px;
  border-left: 2px solid #e8e8e8;
  margin-left: 20px;
}

.submenu-item {
  padding: 10px 16px;
  margin: 2px 0;
  border-radius: 6px;
  text-decoration: none;
  color: #888;
  font-size: 13px;
  transition: all 0.2s ease;
}

.submenu-item:hover {
  background: #f0f8f5;
  color: #037266;
}

.submenu-item.active {
  background: #e8f5f3;
  color: #037266;
  font-weight: 500;
}

/* Submenu transition */
.submenu-enter-active,
.submenu-leave-active {
  transition: all 0.3s ease;
  overflow: hidden;
}

.submenu-enter-from,
.submenu-leave-to {
  opacity: 0;
  max-height: 0;
}

.submenu-enter-to,
.submenu-leave-from {
  opacity: 1;
  max-height: 500px;
}

/* Topbar */
.topbar {
  display: flex;
  align-items: center;
  padding: 16px 24px;
  background: #fefdf7;
  border-bottom: 1px solid #e8e8e8;
}

.hamburger {
  background: none;
  border: none;
  cursor: pointer;
  margin-right: 16px;
  padding: 8px;
  border-radius: 6px;
  transition: background 0.2s ease;
}

.hamburger:hover {
  background: #f0f8f5;
}

.hamburger-icon {
  width: 24px;
  height: 24px;
  color: #037266;
}

/* Breadcrumb Container */
.breadcrumb-container {
  flex: 1;
}

.breadcrumb-line {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 15px;
}

.breadcrumb-title {
  font-size: 18px;
  font-weight: 600;
  color: #037266;
  white-space: nowrap;
}

.breadcrumb-divider {
  color: #d0d0d0;
  font-weight: 300;
  font-size: 18px;
}

.breadcrumb-path {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: #888;
}

.breadcrumb-link {
  color: #888;
  text-decoration: none;
  transition: color 0.2s ease;
}

.breadcrumb-link:hover {
  color: #037266;
}

.breadcrumb-arrow {
  color: #d0d0d0;
  font-size: 12px;
  margin: 0 2px;
}

.breadcrumb-current {
  color: #037266;
  font-weight: 500;
}

/* Main */
.main-wrapper {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

/* Content Wrapper - White Background Area */
.content-wrapper {
  flex: 1;
  padding: 20px 24px;
  background: #f5f5f5;
  overflow-y: auto;
}

.main {
  background: #ffffff;
  border-radius: 12px;
  padding: 24px;
  min-height: calc(100vh - 120px);
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.06);
}

/* Responsive */
@media (max-width: 768px) {
  .sidebar {
    position: fixed;
    left: 0;
    top: 0;
    height: 100vh;
    z-index: 100;
    transform: translateX(0);
  }

  .sidebar.collapsed {
    transform: translateX(-100%);
  }

  .breadcrumb-title {
    font-size: 16px;
  }

  .breadcrumb-path {
    font-size: 12px;
  }

  .breadcrumb-line {
    flex-wrap: wrap;
    gap: 8px;
  }
}
</style>