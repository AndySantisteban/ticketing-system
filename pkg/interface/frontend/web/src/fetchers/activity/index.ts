import axios from "axios";
import { Activity } from "../../models/api-models";

export const getActivitiesByID = async (orderId: number) => {
  try {
    const response = await axios.get("/api/Activity/List", {
      params: {
        id: orderId,
        limit: 100,
        offset: 0,
      },
      headers: {
        Authorization: "Bearer " + localStorage.getItem("@token-infositel"),
      },
    });
    return response.data;
  } catch (error) {
    throw new Error("Error fetching users");
  }
};

export const createActivity = async (activity: Activity) => {
  try {
    const response = await axios.post("/api/Activity/Create", activity, {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("@token-infositel"),
      },
    });
    return response.data;
  } catch (error) {
    throw new Error("Error fetching users");
  }
};
