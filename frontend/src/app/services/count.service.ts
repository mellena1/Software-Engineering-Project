import { Injectable } from "@angular/core";
import {
  HttpClient,
  HttpHeaders,
  HttpErrorResponse,
  HttpParams
} from "@angular/common/http";
import { environment } from "../../environments/environment";
import { throwError as observableThrowError } from "rxjs";
import { Count } from "../data_models/count";
import { catchError, map } from "rxjs/operators";
import { WriteResponse } from "./writeResponse";

@Injectable({
  providedIn: "root"
})
export class CountService {
  constructor(private http: HttpClient) {}
  private apiUrl = environment.apiUrl;

  getAllCounts() {
    return this.http.get<Count>(this.apiUrl + "/counts").pipe(
      map(data => data),
      catchError(this.handleError)
    );
  }

  getACount(sessionID: number) {
    var params = new HttpParams().set("id", sessionID.toString());
    return this.http.get<Count[]>(this.apiUrl + "/count", {
      params: params
    });
  }

  getCountsBySpeaker() {
    return this.http.get<Map<String, Map<String, Count[]>>>(
      this.apiUrl + "/countsBySpeaker"
    );
  }

  writeACount(count: Count) {
    return this.http.post<WriteResponse>(this.apiUrl + "/count", count);
  }

  updateCount(count: Count) {
    return this.http.put(this.apiUrl + "/count", count);
  }

  deleteCount(sessionID: number) {
    var params = new HttpParams().set("id", sessionID.toString());
    return this.http.delete(this.apiUrl + "/count", {
      params: params
    });
  }

  private handleError(res: HttpErrorResponse | any) {
    console.error(res.error || res.body.error);
    return observableThrowError(res.error || "Server error");
  }
}
