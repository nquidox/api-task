package repository

func (d *DB) Migrate(payload any) error {
	err := d.DB.AutoMigrate(payload)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) Create(payload any) error {
	err := d.DB.Create(payload).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) ReadAll(payload any) error {
	err := d.DB.Model(payload).Find(payload).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) Update(payload any, userId uint) error {
	err := d.DB.Where("id = ?", userId).Updates(payload).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) Delete(payload any, userId uint) error {
	err := d.DB.Model(payload).Where("id = ?", userId).Delete(payload).Error
	if err != nil {
		return err
	}
	return nil
}
