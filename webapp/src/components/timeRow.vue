<script>
export default {
  data() {
    return {
      user: "a",
      time:"!"
    };
  },
};
</script>

<template>
  <li class="timeRow">
    <span class="user">{{ user }}</span>
    <span class="user">{{ time }}</span>
  </li>
</template>

<style scoped>
li::before {
  position: absolute;
  left: 2rem;
  counter-increment: timers;
  content: counter(timers);
}

li {
  display: flex;
  position: relative;
  justify-content: space-around;

  height: fit-content;
  width: 30rem;

  padding: 1rem;

  box-shadow: 4px 4px 16px 0px rgba(0, 0, 0, 0.15);
  border-radius: 1rem;
  transition: all ease 0.2s;

  background-color: white;
}

li:hover {
  cursor: pointer;
  scale: 1.1;
}

li:marker {
  z-index: 9999;
  color: red;
  font-weight: bold;
}
</style>
