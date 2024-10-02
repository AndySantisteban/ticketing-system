import axios from "axios";
import { CreateClientRouteDTO } from "../../models/api-models";

export const getClients = async () => {
  try {
    const response = await axios.get("/api/Client/List?offset=0&limit=200", {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("@token-infositel"),
      },
    });
    return response.data;
  } catch (error) {
    throw new Error("Error fetching users");
  }
};

export const getClientByID = async (id: number) => {
  try {
    const response = await axios.get("/api/Client/Get", {
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

export const createClient = async (client: CreateClientRouteDTO) => {
  try {
    const response = await axios.post("/api/Client/Create", client, {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("@token-infositel"),
      },
    });
    return response.data;
  } catch (error) {
    throw new Error("Error fetching users");
  }
};
