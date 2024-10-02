export type UuidType = string
export type BigIntType = BigInt
export type DateType = Date


 export interface CreateActivityQuery {
  ID: number
  OrderID: number
  UserID: number
  Date: DateType
  Action: string
  Details: string
}

 export interface Activity {
  ID: number
  OrderID:  {
Int32: number
Valid: boolean
}
  UserID:  {
Int32: number
Valid: boolean
}
  Date:  {
Time: DateType
Valid: boolean
}
  Action:  {
String: string
Valid: boolean
}
  Details:  {
String: string
Valid: boolean
}
}

 export interface GetActivityByUidQuery {
  Id: number
}

 export interface Activity {
  ID: number
  OrderID:  {
Int32: number
Valid: boolean
}
  UserID:  {
Int32: number
Valid: boolean
}
  Date:  {
Time: DateType
Valid: boolean
}
  Action:  {
String: string
Valid: boolean
}
  Details:  {
String: string
Valid: boolean
}
}

 export interface ListActivityByOrderIDQuery {
  Offset: number
  Limit: number
  Id: number
}

 export type ActivityArray = Activity[] 

export interface Activity {
  ID: number
  OrderID:  {
Int32: number
Valid: boolean
}
  UserID:  {
Int32: number
Valid: boolean
}
  Date:  {
Time: DateType
Valid: boolean
}
  Action:  {
String: string
Valid: boolean
}
  Details:  {
String: string
Valid: boolean
}
}

 export interface Client {
  ID: number
  Name: string
  Address:  {
String: string
Valid: boolean
}
  District:  {
String: string
Valid: boolean
}
  City:  {
String: string
Valid: boolean
}
  Country:  {
String: string
Valid: boolean
}
  Phone:  {
String: string
Valid: boolean
}
  Ruc:  {
String: string
Valid: boolean
}
  ContactPerson:  {
String: string
Valid: boolean
}
  Email:  {
String: string
Valid: boolean
}
  Website:  {
String: string
Valid: boolean
}
  AddressLine2:  {
String: string
Valid: boolean
}
  PostalCode:  {
String: string
Valid: boolean
}
  Fax:  {
String: string
Valid: boolean
}
  Notes:  {
String: string
Valid: boolean
}
}

 export interface GetClientByIDQuery {
  Id: number
}

 export interface ListAllClientsQuery {
  Offset: number
  Limit: number
}

 export type ClientArray = Client[] 

export interface Client {
  ID: number
  Name: string
  Address:  {
String: string
Valid: boolean
}
  District:  {
String: string
Valid: boolean
}
  City:  {
String: string
Valid: boolean
}
  Country:  {
String: string
Valid: boolean
}
  Phone:  {
String: string
Valid: boolean
}
  Ruc:  {
String: string
Valid: boolean
}
  ContactPerson:  {
String: string
Valid: boolean
}
  Email:  {
String: string
Valid: boolean
}
  Website:  {
String: string
Valid: boolean
}
  AddressLine2:  {
String: string
Valid: boolean
}
  PostalCode:  {
String: string
Valid: boolean
}
  Fax:  {
String: string
Valid: boolean
}
  Notes:  {
String: string
Valid: boolean
}
}

 export interface CreateClientQuery {
  ID: number
  Name: string
  Address: string
  District: string
  City: string
  Country: string
  Phone: string
  Ruc: string
  ContactPerson: string
  Email: string
  Website: string
  AddressLine2: string
  PostalCode: string
  Fax: string
  Notes: string
}

 export interface UpdateClientQuery {
  ID: number
  Name: string
  Address: string
  District: string
  City: string
  Country: string
  Phone: string
  Ruc: string
  ContactPerson: string
  Email: string
  Website: string
  AddressLine2: string
  PostalCode: string
  Fax: string
  Notes: string
}

 export interface DeleteClientByIDQuery {
  Id: number
}

 export interface GetCommentByIDQuery {
  Id: number
}

 export interface ListCommentsByOrderIDQuery {
  Offset: number
  Limit: number
  OrderID: number
}

 export interface Comment {
  ID: number
  OrderID:  {
Int32: number
Valid: boolean
}
  UserID:  {
Int32: number
Valid: boolean
}
  Date:  {
Time: DateType
Valid: boolean
}
  Comment:  {
String: string
Valid: boolean
}
}

 export type CommentArray = Comment[] 

export interface Comment {
  ID: number
  OrderID:  {
Int32: number
Valid: boolean
}
  UserID:  {
Int32: number
Valid: boolean
}
  Date:  {
Time: DateType
Valid: boolean
}
  Comment:  {
String: string
Valid: boolean
}
}

 export interface UpdateCommentQuery {
  ID: number
  OrderID: number
  UserID: number
  Date: DateType
  Comment: string
}

 export interface CreateCommentQuery {
  ID: number
  OrderID: number
  UserID: number
  Date: DateType
  Comment: string
}

 export interface DeleteCommentQuery {
  Id: number
}

 export interface GetOrderStatusByIDQuery {
  ID: number
}

 export interface OrderStatus {
  ID: number
  Name: string
}

 export interface ListOrderStatusQuery {
  Limit: number
  Offset: number
}

 export type OrderStatusArray = OrderStatus[] 

export interface OrderStatus {
  ID: number
  Name: string
}

 export interface CreateOrderStatusQuery {
  ID: number
  Name: string
}

 export interface OrderStatus {
  ID: number
  Name: string
}

 export interface UpdateOrderStatusQuery {
  ID: number
  Name: string
}

 export interface GetOrderQuery {
  Id: number
}

 export interface ListOrdersQuery {
  Offset: number
  Limit: number
}

 export type OrderArray = Order[] 

export interface Order {
  ID: number
  ClientID:  {
Int32: number
Valid: boolean
}
  EquipmentID:  {
Int32: number
Valid: boolean
}
  OrderNumber: string
  ReportedIssue:  {
String: string
Valid: boolean
}
  Diagnosis:  {
String: string
Valid: boolean
}
  Solution:  {
String: string
Valid: boolean
}
  EstimatedTime:  {
Int64: BigIntType
Valid: boolean
}
  Budget:  {
String: string
Valid: boolean
}
  StatusID:  {
Int32: number
Valid: boolean
}
  AssignedTo:  {
Int32: number
Valid: boolean
}
  CreationDate:  {
Time: DateType
Valid: boolean
}
  Priority:  {
String: string
Valid: boolean
}
}

 export interface Order {
  ID: number
  ClientID:  {
Int32: number
Valid: boolean
}
  EquipmentID:  {
Int32: number
Valid: boolean
}
  OrderNumber: string
  ReportedIssue:  {
String: string
Valid: boolean
}
  Diagnosis:  {
String: string
Valid: boolean
}
  Solution:  {
String: string
Valid: boolean
}
  EstimatedTime:  {
Int64: BigIntType
Valid: boolean
}
  Budget:  {
String: string
Valid: boolean
}
  StatusID:  {
Int32: number
Valid: boolean
}
  AssignedTo:  {
Int32: number
Valid: boolean
}
  CreationDate:  {
Time: DateType
Valid: boolean
}
  Priority:  {
String: string
Valid: boolean
}
}

 export interface CreateOrderQuery {
  ID: number
  ClientID: number
  Equipement: number
  OrderNumber: string
  ReportedIssue: string
  Diagnosis: string
  Solution: string
  EstimatedTime: BigIntType
  Budget: string
  StatusID: number
  AssignedTo: number
  CreationDate: DateType
  Priority: string
}

 export interface CreateRouteDTO {
  id : number
  statusID : number
  assignedTo : number
  creationDate : DateType
  clientID : number
  priority : string
  equipement : number
  orderNumber : string
  reportedIssue : string
  diagnosis : string
  solution : string
  estimatedTime : BigIntType
  budget : string
}

 export interface Order {
  ID: number
  ClientID:  {
Int32: number
Valid: boolean
}
  EquipmentID:  {
Int32: number
Valid: boolean
}
  OrderNumber: string
  ReportedIssue:  {
String: string
Valid: boolean
}
  Diagnosis:  {
String: string
Valid: boolean
}
  Solution:  {
String: string
Valid: boolean
}
  EstimatedTime:  {
Int64: BigIntType
Valid: boolean
}
  Budget:  {
String: string
Valid: boolean
}
  StatusID:  {
Int32: number
Valid: boolean
}
  AssignedTo:  {
Int32: number
Valid: boolean
}
  CreationDate:  {
Time: DateType
Valid: boolean
}
  Priority:  {
String: string
Valid: boolean
}
}

 export interface UpdateOrderQuery {
  ID: number
  ClientID: number
  Equipement: number
  OrderNumber: string
  ReportedIssue: string
  Diagnosis: string
  Solution: string
  EstimatedTime: BigIntType
  Budget: string
  StatusID: number
  AssignedTo: number
  CreationDate: DateType
  Priority: string
}

 export interface UpdateRouteDTO {
  id : number
  statusID : number
  assignedTo : number
  creationDate : DateType
  clientID : number
  priority : string
  equipement : number
  orderNumber : string
  reportedIssue : string
  diagnosis : string
  solution : string
  estimatedTime : BigIntType
  budget : string
}

 export interface GetUserByIDQuery {
  ID: number
}

 export interface User {
  ID: number
  Name: string
  Email: string
  PermissionType: string
  CreationDate:  {
Time: DateType
Valid: boolean
}
  InactiveStatus:  {
String: string
Valid: boolean
}
  Password:  {
String: string
Valid: boolean
}
}

 export interface ListUserQuery {
  Limit: number
  Offset: number
}

 export type UserArray = User[] 

export interface User {
  ID: number
  Name: string
  Email: string
  PermissionType: string
  CreationDate:  {
Time: DateType
Valid: boolean
}
  InactiveStatus:  {
String: string
Valid: boolean
}
  Password:  {
String: string
Valid: boolean
}
}

 export interface CreateClientRouteDTO {
  id : number
  name : string
  address : string
  district : string
  city : string
  country : string
  phone : string
  ruc : string
  contactPerson : string
  email : string
  website : string
  addressLine2 : string
  postalCode : string
  fax : string
  notes : string
}

 export interface CreateUserDTO {
  id : number
  name : string
  email : string
  permission_type : string
  creation_date : DateType
  inactive_status : string
  password : string
}

 export interface CreateActivityQuery {
  id : number
  orderID? : number
  userID? : number
  date? : DateType
  action? : string
  details? : string
}

 export interface ListActivityByOrderIDRouteDTO {
  Offset: number
  Limit: number
  Id: number
}

 export interface GetActivityByUidRouteDTO {
  Id: number
}

 export interface Equipment {
  ID: number
  TypeID:  {
Int32: number
Valid: boolean
}
  Name: string
  SerialNumber: string
  Notes:  {
String: string
Valid: boolean
}
}

 export type EquipmentArray = Equipment[] 

export interface Equipment {
  ID: number
  TypeID:  {
Int32: number
Valid: boolean
}
  Name: string
  SerialNumber: string
  Notes:  {
String: string
Valid: boolean
}
}

 export interface GetEquipmentQueryDTO {
  ID: number
}

 export interface ListAllEquipmentQueryDTO {
  Offset: number
  Limit: number
}

 export interface CreateEquipmentQueryDTO {
  id? : number
  type_id? : number
  name? : string
  serial_number? : string
  notes? : string
}

 export interface DeleteEquipmentTypeQueryDTO {
  Id: number
}

 export type EquipmentTypeArray = EquipmentType[] 

export interface EquipmentType {
  ID: number
  Name: string
}

 export interface EquipmentType {
  ID: number
  Name: string
}

 export interface GetEquipmentTypeQueryDTO {
  ID: number
}

 export interface ListAllEquipmentTypeQueryDTO {
  Offset: number
  Limit: number
}

 export interface CreateEquipmentTypeQueryDTO {
  id? : number
  name? : string
}

 export interface DeleteEquipmentTypeRouteDTO {
  ID: number
}



