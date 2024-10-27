import { baseUrl, jsonHeaders } from "./config";

export async function setRemoteVolume(volume: number): Promise<void> {
  console.log(`Attempting to set remote volume to ${volume}`);

  const reqData = {volume};

  await fetch(baseUrl + "/api/playback/volume", {
    method: "POST",
    body: JSON.stringify(reqData),
    headers: jsonHeaders
  });
}