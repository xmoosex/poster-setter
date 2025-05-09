import apiClient from "./apiClient";
import { APIResponse } from "../types/apiResponse";
import { LibrarySection, MediaItem } from "../types/mediaItem";
import { ClientMessage } from "../types/clientMessage";
import { ReturnErrorMessage } from "./api.shared";

export const fetchMediaServerLibraryItems = async (): Promise<
	APIResponse<LibrarySection[]>
> => {
	try {
		const response = await apiClient.get<APIResponse<LibrarySection[]>>(
			`/mediaserver/sections/all/`
		);
		return response.data;
	} catch (error) {
		return ReturnErrorMessage<LibrarySection[]>(error);
	}
};

export const fetchMediaServerItemContent = async (
	ratingKey: string
): Promise<APIResponse<MediaItem>> => {
	try {
		const response = await apiClient.get<APIResponse<MediaItem>>(
			`/mediaserver/item/${ratingKey}`
		);
		return response.data;
	} catch (error) {
		return ReturnErrorMessage<MediaItem>(error);
	}
};

export const postSendSetToAPI = async (
	sendData: ClientMessage
): Promise<APIResponse<null>> => {
	try {
		const response = await apiClient.post<APIResponse<null>>(
			`/mediaserver/update/send`,
			sendData
		);
		return response.data;
	} catch (error) {
		return ReturnErrorMessage<null>(error);
	}
};

export const fetchMediaServerImageData = async (
	ratingKey: string,
	type: string
): Promise<string> => {
	try {
		const API_URL = `/mediaserver/image/${ratingKey}/${type}`;
		const response = await apiClient.get<APIResponse<null>>(API_URL);
		if (response.status !== 200) {
			throw new Error("Failed to fetch image data");
		}
		return "/api" + API_URL;
	} catch {
		return "/logo.png"; // Fallback image
	}
};
