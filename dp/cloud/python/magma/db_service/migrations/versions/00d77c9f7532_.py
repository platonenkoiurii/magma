"""empty message

Revision ID: 00d77c9f7532
Revises: e793671a19a6
Create Date: 2022-05-04 17:04:29.237458

"""
import sqlalchemy as sa
from alembic import op

# revision identifiers, used by Alembic.
revision = '00d77c9f7532'
down_revision = 'e793671a19a6'
branch_labels = None
depends_on = None


def upgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.add_column('cbsds', sa.Column('single_step_enabled', sa.Boolean(), server_default='false', nullable=False))
    op.add_column('cbsds', sa.Column('cbsd_category', sa.String(), server_default='B', nullable=False))
    op.add_column('cbsds', sa.Column('latitude_deg', sa.Float(), nullable=True))
    op.add_column('cbsds', sa.Column('longitude_deg', sa.Float(), nullable=True))
    op.add_column('cbsds', sa.Column('height_m', sa.Float(), nullable=True))
    op.add_column('cbsds', sa.Column('height_type', sa.String(), nullable=True))
    op.add_column('cbsds', sa.Column('horizontal_accuracy_m', sa.Float(), nullable=True))
    op.add_column('cbsds', sa.Column('antenna_azimuth_deg', sa.Integer(), nullable=True))
    op.add_column('cbsds', sa.Column('antenna_downtilt_deg', sa.Integer(), nullable=True))
    op.add_column('cbsds', sa.Column('antenna_beamwidth_deg', sa.Integer(), nullable=True))
    op.add_column('cbsds', sa.Column('antenna_model', sa.String(), nullable=True))
    op.add_column('cbsds', sa.Column('eirp_capability_dbm_mhz', sa.Integer(), nullable=True))
    op.add_column('cbsds', sa.Column('cpi_digital_signature', sa.Text(), nullable=True))
    op.add_column('cbsds', sa.Column('indoor_deployment', sa.Boolean(), server_default='false', nullable=False))
    # ### end Alembic commands ###


def downgrade():
    # ### commands auto generated by Alembic - please adjust! ###
    op.drop_column('cbsds', 'indoor_deployment')
    op.drop_column('cbsds', 'cpi_digital_signature')
    op.drop_column('cbsds', 'eirp_capability_dbm_mhz')
    op.drop_column('cbsds', 'antenna_model')
    op.drop_column('cbsds', 'antenna_beamwidth_deg')
    op.drop_column('cbsds', 'antenna_downtilt_deg')
    op.drop_column('cbsds', 'antenna_azimuth_deg')
    op.drop_column('cbsds', 'horizontal_accuracy_m')
    op.drop_column('cbsds', 'height_type')
    op.drop_column('cbsds', 'height_m')
    op.drop_column('cbsds', 'longitude_deg')
    op.drop_column('cbsds', 'latitude_deg')
    op.drop_column('cbsds', 'cbsd_category')
    op.drop_column('cbsds', 'single_step_enabled')
    # ### end Alembic commands ###
