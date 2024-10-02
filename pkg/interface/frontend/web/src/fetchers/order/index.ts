import axios from "axios";
import { CreateRouteDTO, UpdateRouteDTO } from "../../models/api-models";

export const createOrder = async (data: CreateRouteDTO) => {
  try {
    const panels = await axios.post("/api/Order/Create", data, {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("@token-infositel"),
      },
    });

    return panels.data;
  } catch (e) {
    throw new Error("Error fetching users");
  }
};

export const updateOrder = async (data: UpdateRouteDTO) => {
  try {
    const panels = await axios.put("/api/Order/Update", data, {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("@token-infositel"),
      },
    });

    return panels.data;
  } catch (e) {
    throw new Error("Error fetching users");
  }
};

export const getOrders = async () => {
  try {
    const panels = await axios.get("/api/Order/List?offset=0&limit=10000", {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("@token-infositel"),
      },
    });

    return panels.data;
  } catch (e) {
    throw new Error("Error fetching users");
  }
};

export const getOrdersByID = async (uid: number) => {
  try {
    const panels = await axios.get("/api/Order/Get", {
      params: {
        id: uid,
      },
      headers: {
        Authorization: "Bearer " + localStorage.getItem("@token-infositel"),
      },
    });

    return panels.data;
  } catch (e) {
    throw new Error("Error fetching orders by ID");
  }
};
