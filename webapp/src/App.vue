<script setup>
import TimeRow from './components/timeRow.vue';
import NewsTable from './components/NewsTable.vue';

</script>


<template>
  <div class="mainContent">
    <Suspense>

      <NewsTable/>
      <template #fallback>
        Loading...
      </template>
    </Suspense>
    <!-- <ol type="1">
            <TimeRow user="rgnwld" time="10:00.00" />
            <TimeRow user="rgnwld" time="10:00.00" />
            <TimeRow user="rgnwld" time="10:00.00" />
            <TimeRow user="rgnwld" time="10:00.00" />
            <TimeRow user="rgnwld" time="10:00.00" />
            <TimeRow user="rgnwld" time="10:00.00" />
            <TimeRow user="rgnwld" time="10:00.00" />
            <TimeRow user="rgnwld" time="10:00.00" />
          </ol> -->
  </div>
</template>

<style scoped>
.mainContent {
  display: flex;
  height: 100vh;
  width: 100vw;
  justify-content: center;
}

.mainContent>ol {
  counter-reset: timers;
  display: flex;

  justify-content: center;
  align-items: center;

  flex-direction: column;

  gap: .5rem;
  list-style: decimal;
}
</style>
