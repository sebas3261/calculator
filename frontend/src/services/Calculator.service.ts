import { http } from "../lib/http";
import type {
  CalculationRequest,
  CalculationResponse,
} from "../models/calculator.model";

export const calculatorService = {
  calculate: (expression: string) =>
    http.post<CalculationResponse>("/calculate", {
      expression,
    } satisfies CalculationRequest),
};
