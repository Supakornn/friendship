<script>
  import axios from "axios";

  let message = "";
  let from = "";

  const createFriendShip = async () => {
    try {
      await axios.post("http://localhost:8000/create", {
        message: message,
        from: from
      });
      console.log("Success");
    } catch (error) {
      console.error(error);
    }
  };
</script>

<body
  class="bg-[#88c6a5] min-h-[70vh] box-border flex justify-center items-center text-[hsl(198,1%,29%)] text-center text-[130%] p-4 sm:min-h-screen"
>
  <section
    class="w-full h-auto min-h-[500px] max-w-[500px] min-w-[250px] bg-[25px_25px] shadow-[4px_3px_7px_2px_#00000040] box-border p-4 rounded-[20px] bg-[#f1f5f8]"
  >
    <div class="flex items-center justify-center mb-4">
      <img
        class="w-[24%]"
        src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/756881/laptop.svg"
        alt="pic"
      />
      <h1
        class="rotate-2 bg-[#ffeead] text-2xl px-[1.2rem] py-[0.2rem] rounded-[20%_5%_20%_5%/5%_20%_25%_20%]"
      >
        Friendship
      </h1>
    </div>
    <form on:submit={createFriendShip}>
      <p class="block mb-2 text-2xl">~ To... Gap ~</p>
      <textarea
        class=" box-border bg-transparent text-base text-[rgba(63,62,65,0.7)] w-[90%] mb-2 p-[0.7rem] rounded-br-[15px_3px] rounded-bl-[3px_15px] border-dashed border-2 border-[#ffcc5c] focus:outline-none
      "
        id="message"
        name="message"
        bind:value={message}
        placeholder="เขียนข้อความ......"
        required
      />

      <p class="text-left ml-[60px]">~ From ...</p>
      <input
        class=" box-border bg-transparent text-base text-[rgba(63,62,65,0.7)] w-[40%] mr-5 p-[0.7rem] rounded-br-[15px_3px] rounded-bl-[3px_15px] border-dashed border-2 border-[#ffcc5c] focus:outline-none"
        id="from"
        name="from"
        bind:value={from}
        placeholder="ส่งมาจากใคร?"
        required
      />

      <button
        class="rotate-[4deg] origin-center no-underline shadow-[0_2px_0_#494a4b] transition-all duration-[0.3s] ease-[cubic-bezier(0.175,0.885,0.32,1.275)] bg-[url('data:image/gif;base64,R0lGODlhBAAEAIABAAAAAAAAACH/C1hNUCBEYXRhWE1QPD94cGFja2V0IGJlZ2luPSLvu78iIGlkPSJXNU0wTXBDZWhpSHpyZVN6TlRjemtjOWQiPz4gPHg6eG1wbWV0YSB4bWxuczp4PSJhZG9iZTpuczptZXRhLyIgeDp4bXB0az0iQWRvYmUgWE1QIENvcmUgNS4wLWMwNjEgNjQuMTQwOTQ5LCAyMDEwLzEyLzA3LTEwOjU3OjAxICAgICAgICAiPiA8cmRmOlJERiB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiPiA8cmRmOkRlc2NyaXB0aW9uIHJkZjphYm91dD0iIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHhtbG5zOnhtcE1NPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvbW0vIiB4bWxuczpzdFJlZj0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL3NUeXBlL1Jlc291cmNlUmVmIyIgeG1wOkNyZWF0b3JUb29sPSJBZG9iZSBQaG90b3Nob3AgQ1M1LjEgV2luZG93cyIgeG1wTU06SW5zdGFuY2VJRD0ieG1wLmlpZDo5NUY1OENCRDdDMDYxMUUyOTEzMEE1MEM5QzM0NDVBMyIgeG1wTU06RG9jdW1lbnRJRD0ieG1wLmRpZDo5NUY1OENCRTdDMDYxMUUyOTEzMEE1MEM5QzM0NDVBMyI+IDx4bXBNTTpEZXJpdmVkRnJvbSBzdFJlZjppbnN0YW5jZUlEPSJ4bXAuaWlkOjk1RjU4Q0JCN0MwNjExRTI5MTMwQTUwQzlDMzQ0NUEzIiBzdFJlZjpkb2N1bWVudElEPSJ4bXAuZGlkOjk1RjU4Q0JDN0MwNjExRTI5MTMwQTUwQzlDMzQ0NUEzIi8+IDwvcmRmOkRlc2NyaXB0aW9uPiA8L3JkZjpSREY+IDwveDp4bXBtZXRhPiA8P3hwYWNrZXQgZW5kPSJyIj8+Af/+/fz7+vn49/b19PPy8fDv7u3s6+rp6Ofm5eTj4uHg397d3Nva2djX1tXU09LR0M/OzczLysnIx8bFxMPCwcC/vr28u7q5uLe2tbSzsrGwr66trKuqqainpqWko6KhoJ+enZybmpmYl5aVlJOSkZCPjo2Mi4qJiIeGhYSDgoGAf359fHt6eXh3dnV0c3JxcG9ubWxramloZ2ZlZGNiYWBfXl1cW1pZWFdWVVRTUlFQT05NTEtKSUhHRkVEQ0JBQD8+PTw7Ojk4NzY1NDMyMTAvLi0sKyopKCcmJSQjIiEgHx4dHBsaGRgXFhUUExIREA8ODQwLCgkIBwYFBAMCAQAAIfkEAQAAAQAsAAAAAAQABAAAAgYEEpdoeQUAOw==')]
         bg-[#ff6f69] pb-[3px] p-0 rounded-[5px] border-[none] focus:translate-y-1 focus:pb-0 focus:outline-0 active:translate-y-1 active:pb-0 active:outline-0"
        type="submit"
        ><span
          class="block px-4 py-2 rounded-[5px] border-2 border-solid border-[#494a4b] bg-[#f1f5f8]"
          >ส่งFriendship</span
        ></button
      >
    </form>

    <div />
  </section>
</body>
