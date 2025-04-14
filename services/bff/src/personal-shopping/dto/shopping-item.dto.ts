import { Field, ObjectType, registerEnumType } from '@nestjs/graphql'

@ObjectType()
export class ShoppingItem {
  @Field()
  id: number

  @Field()
  name: string

  @Field()
  status: Status

  @Field()
  categoryId: number
}

export enum Status {
  UNPURCHASED = 'UNPURCHASED',
  PURCHASED = 'PURCHASED',
}

registerEnumType(Status, { name: 'PersonalShoppingStatus' })
