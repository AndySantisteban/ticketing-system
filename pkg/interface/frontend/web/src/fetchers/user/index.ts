import axios from "axios";
import { CreateUserDTO } from "../../models/api-models";

export const getUsers = async () => {
  try {
    const response = await axios.get("/api/User/List?offset=0&limit=200", {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("@token-infositel"),
      },
    });
    return response.data;
  } catch (error) {
    throw new Error("Error fetching users");
  }
};

export const getUserByID = async (id: number) => {
  try {
    const response = await axios.get("/api/User/Get", {
      params: {
        id,
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

export const createUsers = async (user: CreateUserDTO) => {
  try {
    const response = await axios.post("/api/User/Create", user, {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("@token-infositel"),
      },
    });
    return response.data;
  } catch (error) {
    throw new Error("Error fetching users");
  }
};
