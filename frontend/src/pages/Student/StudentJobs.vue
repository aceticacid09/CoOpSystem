<!-- frontend/src/pages/Student/StudentJobs.vue -->
<template>
  <DashboardLayout role="student">
    <JobsSections user-role="student" :show-bookmark="true" :show-actions="true" :initial-bookmarked="[1, 3, 5]"
      :initial-saved="[1, 3, 5]" :initial-applied="[1]" />
  </DashboardLayout>
</template>

<script setup>
import DashboardLayout from '../../components/DashboardLayout.vue';
import JobsSections from "../../components/JobsSections.vue";
</script>

<style scoped></style>