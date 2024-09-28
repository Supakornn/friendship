<script>
  function fetchData() {
    fetch("https://friendship-api.up.railway.app/friendships")
      .then((response) => response.json())
      .then((data) => {
        const container = document.getElementById("container");
        container.innerHTML = "";

        data.forEach((friendship) => {
          const section = document.createElement("section");
          section.classList.add(
            "w-full",
            "h-auto",
            "min-h-[50px]",
            "max-w-[500px]",
            "min-w-[500px]",
            "bg-[25px_25px]",
            "shadow-[4px_3px_7px_2px_#00000040]",
            "box-border",
            "p-4",
            "rounded-[20px]",
            "bg-[#f1f5f8]"
          );
          section.style.display = "flex";
          section.style.flexDirection = "row";
          section.style.justifyContent = "center";
          section.style.alignItems = "center";

          const contentContainer = document.createElement("div");
          contentContainer.style.marginLeft = "1rem";
          contentContainer.style.display = "flex";
          contentContainer.style.flexDirection = "column";
          contentContainer.style.gap = "0.5rem";

          const h1 = document.createElement("h1");
          h1.classList.add(
            "rotate-2",
            "bg-[#ffeead]",
            "text-2xl",
            "px-[1.2rem]",
            "py-[0.2rem]",
            "rounded-[20%_5%_20%_5%/5%_20%_25%_20%]"
          );
          h1.textContent = "Friendship";

          const pMessage = document.createElement("p");
          pMessage.textContent = friendship.message;
          pMessage.style.fontSize = "25px";
          pMessage.style.fontWeight = "bold";

          const pFrom = document.createElement("p");
          pFrom.textContent = `จาก: ${friendship.from}`;
          pFrom.style.marginLeft = "300px";
          pFrom.style.marginTop = "auto";

          contentContainer.appendChild(h1);
          contentContainer.appendChild(pMessage);
          contentContainer.appendChild(pFrom);

          section.appendChild(contentContainer);
          container.appendChild(section);
        });
      })
      .catch((error) => {
        console.log(error);
      });
  }

  fetchData();
  setInterval(fetchData, 5000);
</script>

<body
  class="bg-[#3a534b] h-screen box-border flex justify-center items-center text-[hsl(198,1%,29%)] text-center text-[130%]"
>
  <div id="container" class="flex flex-row flex-wrap justify-center" />

  <style>
    #container > section {
      margin: 10px;
    }
  </style>
</body>
