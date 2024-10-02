import axios from "axios";
import { CreateCommentQuery } from "../../models/api-models";

export const getComment = async (orderId: number) => {
  try {
    const response = await axios.get("/api/Comment/List", {
      params: {
        orderID: orderId,
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

export const createComment = async (comment: CreateCommentQuery) => {
  try {
    const response = await axios.post("/api/Comment/Create", comment, {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("@token-infositel"),
      },
    });
    return response.data;
  } catch (error) {
    throw new Error("Error fetching users");
  }
};
