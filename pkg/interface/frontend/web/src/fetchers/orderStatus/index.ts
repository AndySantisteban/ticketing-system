import axios from "axios";

export const getPanels = async () => {
  try {
    const panels = await axios.get("/api/OrderStatus/List?offset=0&limit=20", {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("@token-infositel"),
      },
    });

    return panels.data;
  } catch (e) {
    throw new Error("Error fetching users");
  }
};
