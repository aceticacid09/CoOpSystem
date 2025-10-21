<!-- frontend/src/pages/Teacher/TeacherNewsEvents.vue -->
<template>
  <DashboardLayout role="teacher">
    <NewsSections 
      :show-bookmark="true"
      :initial-bookmarked="[1, 2]"
      :initial-saved="[1]"
    />
  </DashboardLayout>
</template>

<script setup>
import DashboardLayout from '../../components/DashboardLayout.vue';
import NewsSections from "../../components/NewsSections.vue";
</script>

<style scoped></style>
