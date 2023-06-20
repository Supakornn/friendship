<script>
  import axios from "axios";
  const createTodo = async () => {
    try {
      await axios.post("http://localhost:8000/api/todos", {
        title: "Test"
      });
      console.log("Success");
    } catch (error) {
      console.error(error);
    }
  };
</script>

<h1>Test</h1>
<button on:click={createTodo}>Create Todo</button>
