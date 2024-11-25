package store_test

//func TestItemsRepo_CreateItems(t *testing.T) {
//	dataBase := build.NewStore()
//	defer dataBase.Close()
//	type fields struct {
//		client *bun.DB
//	}
//	type args struct {
//		ctx   context.Context
//		items *repomodels.Items
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		{
//			name: "create items test",
//			fields: fields{
//				client: dataBase,
//			},
//			args: args{
//				ctx: context.Background(),
//				items: &repomodels.Items{
//					ItemName:    "затраты",
//					GUID:        uuid.New(),
//					Description: "тестик",
//				},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			i := &store.ItemsRepo{
//				client: tt.fields.client,
//			}
//			if err := i.CreateItems(tt.args.ctx, tt.args.items); (err != nil) != tt.wantErr {
//				t.Errorf("CreateItems() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

//func TestItemsRepo_DeleteItems(t *testing.T) {
//	type fields struct {
//		client *bun.DB
//	}
//	type args struct {
//		ctx context.Context
//		id  int
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			i := &ItemsRepo{
//				client: tt.fields.client,
//			}
//			if err := i.DeleteItems(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
//				t.Errorf("DeleteItems() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestItemsRepo_GetAllItems(t *testing.T) {
//	type fields struct {
//		client *bun.DB
//	}
//	type args struct {
//		ctx context.Context
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    []repomodels.Items
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			i := &ItemsRepo{
//				client: tt.fields.client,
//			}
//			got, err := i.GetAllItems(tt.args.ctx)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetAllItems() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetAllItems() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestItemsRepo_GetOne(t *testing.T) {
//	type fields struct {
//		client *bun.DB
//	}
//	type args struct {
//		ctx    context.Context
//		itemID int
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *repomodels.Items
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			i := &ItemsRepo{
//				client: tt.fields.client,
//			}
//			got, err := i.GetOne(tt.args.ctx, tt.args.itemID)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetOne() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetOne() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestItemsRepo_UpdateItems(t *testing.T) {
//	type fields struct {
//		client *bun.DB
//	}
//	type args struct {
//		ctx context.Context
//		u   *repomodels.Items
//		id  int
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			i := &ItemsRepo{
//				client: tt.fields.client,
//			}
//			if err := i.UpdateItems(tt.args.ctx, tt.args.u, tt.args.id); (err != nil) != tt.wantErr {
//				t.Errorf("UpdateItems() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestNewItemsRepo(t *testing.T) {
//	type args struct {
//		client *bun.DB
//	}
//	tests := []struct {
//		name string
//		args args
//		want *ItemsRepo
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewItemsRepo(tt.args.client); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewItemsRepo() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
