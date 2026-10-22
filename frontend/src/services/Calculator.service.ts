import { http } from "../lib/http";
import type { CalculationRequest } from "../models/CalculationRequest";
import type { CalculationResponse } from "../models/CalculationResponse";

export const calculatorService = {
  calculate: (expression: string) =>
    http.post<CalculationResponse>("/calculate", {
      expression,
    } satisfies CalculationRequest),
};
