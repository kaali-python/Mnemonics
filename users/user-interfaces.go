
package users

import ("time"
  _ "github.com/satori/go.uuid"


)



type Response struct{
    Message string `json:"message"`
    Error bool `json:"error"`
    Success bool `json:"success"`
    Data interface{} `json:"data"`
    Token string `json:"token"`
}

type KeyRequest struct{
    Email string `json:"email"`
    PhoneNumber string `bson:"phone_number" json:"phone_number"`
}




type GenerateMnemonic struct{
    Mnemonic string `json:"mnemonic"`
    MasterPublicKey string `json:"master_public_key"`
    MasterPrivateKey string `json:"master_private_key"`
    //SerializedPublicKey string `gorethink:"serialized_public_key" bson:"serialized_public_key" json:"serialized_public_key"`

    ZerothPublicKey string `json:"zeroth_public_key"`
    ZerothPrivateKey string `json:"zeroth_private_key"`
    //SerializedZerothPublicKey string `gorethink:"serialized_zeroth_public_key" bson:"serialized_zeroth_public_key" json:"serialized_zeroth_public_key"`


}




type UserStruct struct{
    UserID string `bson:"user_id" json:"user_id"`
    Address string `bson:"address" json:"address"`
    Email string `bson:"email" json:"email"`
    Adhaar string `bson:"adhaar" json:"adhaar"`
    FirstName string `bson:"first_name" json:"first_name"`
    LastName string `bson:"last_name" json:"last_name"`
    CreatedAt time.Time `bson:"created_at" json:"created_at"`
    PhoneNumber string `bson:"phone_number" json:"phone_number"`
    PanCard string `bson:"pancard" json:"pancard"`
    Details map[string]string `json:"details"`
    PublicKey string `gorethink:"public_key" bson:"public_key" json:"public_key"`
    SerializedPublicKey string `gorethink:"serialized_public_key" bson:"serialized_public_key" json:"serialized_public_key"`

    ZerothPublicKey string `gorethink:"zeroth_public_key" bson:"zeroth_public_key" json:"zeroth_public_key"`
    SerializedZerothPublicKey string `gorethink:"serialized_zeroth_public_key" bson:"serialized_zeroth_public_key" json:"serialized_zeroth_public_key"`


}

type SecretsStruct struct {
  UserID string `gorethink:"user_id" bson:"user_id" json:"user_id"`
  SerializedPublicKey string `gorethink:"serialized_public_key" bson:"serialized_public_key" json:"serialized_public_key"`
  PublicKey string `gorethink:"public_key" bson:"public_key" json:"public_key"`
  ZerothPublicKey string `gorethink:"zeroth_public_key" bson:"zeroth_public_key" json:"zeroth_public_key"`
  SerializedZerothPublicKey string `gorethink:"serialized_zeroth_public_key" bson:"serialized_zeroth_public_key" json:"serialized_zeroth_public_key"`
  Secrets []string `gorethink:"secrets" bson:"secrets" json:secrets`
  Adhaar string `gorethink:"adhaar" bson:"adhaar" json:"adhaar"`
  Email string `gorethink:"email" bson:"email" json:"email"`
  CreatedAt time.Time `gorethink:"created_at" bson:"created_at" json:"created_at"`
  PhoneNumber string `gorethink:"phone_number" bson:"phone_number" json:"phone_number"`
  PanCard string `gorethink:"pancard" bson:"pancard" json:"pancard"`
  FirstName string `gorethink:"first_name" bson:"first_name" json:"first_name"`
  LastName string `gorethink:"last_name" bson:"last_name" json:"last_name"`
}

type HSMSecretsStruct struct {
  UseridHash string `bson:"user_id_hash" json:"user_id_hash"`
  AESKey string `bson:"aes_key" json:"aes_key"`
  Secrets []string `bson:"secrets" json:secrets`
}







type Credentials struct{
    UserID *string `bson:"user_id" json:"username"`
    Password *string `json:"password"`
}
